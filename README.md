**setup**
- python,perl and bash needs to be supported
- the ntfs-3g and fuse package need to be installed
- setup rclone with dropbox account globally, root should be able to access it( recommended way is to setup rclone **using** the root account )
- connect drives
- Fill in drive paths in the .env file
- run setup with sudo permissions
- setup cron jobs on root crontab
- adjust frequency of updates using cron jobs


**removing**
- remove cron jobs
- stop and remove docker images ( docker compose down --remove-orphans )
