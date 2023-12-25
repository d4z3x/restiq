This is a simple helper utility for `restic`. 
It lets you specify all your repos in one YAML file:

```yaml
restic_bin: /mnt/cache/restic/restic
repos:
  pics:
    location: /mnt/disks/GRAB_AND_GO_USB_HD/restic/Pictures
    type: none
    compression: max
    max_snapshots: 5
    packsize: 128
    dirs:
      - /mnt/user/Pictures
```
`types` can be whatever restoc supports, i.e. sftp, s3, or none (for local disk)

Passwords and other env variables are read from a file called cred.env, which is just of the form:

```bash
RESTIC_PASSWORD=redacted
AWS_ACCESS_KEY_ID=redacted
AWS_SECRET_ACCESS_KEY=redacted
```

# Usage
```
restiq --help
Configuration based tool for `restic` repos

Usage:
  restiq [command]

Available Commands:
  backup      Backup to Repo
  backupall   Backup All Repos
  build       build info
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Lists Repos
  ls          List files from Snapshot
  repos       Lists Repos
  snapshots   Lists Snapshots for Repo
  stats       Stats on a repo
  tokens      List tokens
  version     version info

Flags:
  -h, --help     help for restiq
  -t, --toggle   Help message for toggle

Use "restiq [command] --help" for more information about a command.
```
# Examples:
```bash
restiq list   # to see repos from YAML (config.yml)
restiq backup # to backup the data (if the repo doesn't exist,
              # it will tell you how to create one with the restic command line)
```
