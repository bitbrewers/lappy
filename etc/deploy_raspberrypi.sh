#!/bin/sh
set -e
set -x

if [ $# -eq 0 ]
then
  echo "define target device (eg. sdb)"
  exit 1
fi

TGTDEV=$1
BOOTPART=$TGTDEV"1"
ROOTPART=$TGTDEV"2"

# Partition SD card
(
echo o # Create a new empty DOS partition table
echo n # Add a new partition
echo p # Primary partition
echo 1 # Partition number
echo   # First sector (Accept default: 1)
echo +100M  # Last sector
echo t # Set partition type
echo c # Set the first partition to type W95 FAT32 (LBA).
echo n # Add a second partition
echo p # Primary partition
echo 2 # Second partition
echo   # First sector (Accept default: 1)
echo   # Last sector (Accept default: last)
echo w # Write changes
) | fdisk $TGTDEV

# Create and mount the FAT filesystem:
mkfs.vfat $BOOTPART
mkdir boot
mount $BOOTPART boot

# Create and mount the ext4 filesystem:
mkfs.ext4 $ROOTPART
mkdir root
mount $ROOTPART root

# Download and extract the root filesystem (as root, not via sudo):
wget http://os.archlinuxarm.org/os/ArchLinuxARM-rpi-3-latest.tar.gz
bsdtar -xpf ArchLinuxARM-rpi-3-latest.tar.gz -C root
sync

# Move boot files to the first partition:
mv root/boot/* boot

# Unmount the two partitions:
umount boot root
