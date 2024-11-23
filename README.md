## Explore

This lightweight, fast, and scalable Network-Attached Storage (NAS) solution is specifically designed for managing home photo media collections. Unlike traditional NAS systems, which can be resource-intensive and require powerful servers to ensure speed, this architecture is capable of running on constrained/minimal hardware settings, including the Raspberry Pi 4. It is fast regardless of the size of the media its hosting and prioritizes photo discovery and exploration—because what’s the point of storing all your photos if you’re never gonna look at them anyway later on? 

The system is designed to provide **99.99995% data durability**. With a **probability of total failure of 0.0000005 per year**, the **MTTDL (Mean Time To Data Loss)** is approximately **2 million years**. [^1]


The system is fully configurable through the `.env` file.

**Drive Selection**:
- Opt for **SSDs/USB drives** for high-speed discovery or **HDDs** for scalable, reliable storage.
- **Recommended Setup**:
    - **Discovery Drive**: A fast USB drive to provide instant access to curated photos.
    - **Hot Drives**: Large-capacity HDDs to maintain always-on data redundancy.
    - **Backup Drives**: HDDs with capacity equal to or larger than hot drives for offsite backups.

**Periodic Data Movements**:

- Define when and how often data flows between drives and configure upload synchronization intervals based on your media size and access patterns for optimal performance.
- **Sample Config:**
    - **Discovery Drive Refresh**: Monthly cycles from hot drives to the discovery drive ensure a fresh and curated set of photos.
    - **Uploads**: Nightly synchronization ensures new uploads are regularly integrated into the system.
    - **Full Backups**: Annual backups from the hot drives to an offsite drive protect the entire collection.
    

---

## Setup Instructions

1. **Prerequisites**:
    - Ensure `python`, `perl`, and `bash` are installed and supported.
    - Install `ntfs-3g` and `fuse` packages.
    - Set up `rclone` with a Dropbox account globally, preferably using the root account.
    - Set up docker and docker compose
2. **Environment Configuration**:
    - Connect the necessary drives.
    - Fill in drive paths in the `.env` file.
3. **Running Setup**:
    - Execute the setup script with sudo permissions: `sudo ./setup`
4. **Cron Jobs**:
    - Set up cron jobs in the root crontab to schedule when `updateDISC` (refreshes the discovery drive), `updateHOT` (handles new uploads synchronization onto hot drives and integration into your media library), `backup` (backs up from the hot drives to an offsite drive)
    - Adjust the frequency of updates as required.
    - Optionally you can redirect output from these cronjobs to a log file as a crude form of logging

## Uninstall Instructions

1. **Remove Cron Jobs**:
    - Ensure all cron jobs related to this setup are removed from the root crontab.
2. **Stop and Remove Docker Images (related to homegallery)**:
    - Run the following command: `docker compose down --remove-orphans`

[^1]: Using WD Red drives
