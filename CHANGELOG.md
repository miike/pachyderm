
# Changelog

## 2.0.1
- Fixes a goroutine leak that can occur when a modify file operation is canceled - #7055
- Improves performance by reducing the number of postgres requests needed when renewing chunks being uploaded - #7052
- Improves usability by changing helm defaults- activate enterprise automatically and use default localhost URIs when ingress is disabled - #7039
- Improves memory scaling of join and group inputs beyond millions of datums and - - - - Increases logging verbosity by including etcd and pg-bouncer logs in debug dumps - #7053
## 2.0.0 
Introducing Pachyderm 2.0 with several foundational improvements. Read more details [here](https://www.pachyderm.com/blog/getting-ready-for-pachyderm-2/)

### [What’s new](https://docs.pachyderm.com/2.0.x/getting_started/whats_new/)

- New storage architecture and FileSets for better support for small files, content defined chunking for better de-duplication, automatic compression and encryption of chunks, automatic garbage collection, and more
- Pachyderm Enterprise Management to allow site-wide configuration
- New enterprise UI -- Pachyderm Console
- Improved OAuth-based Authentication
- Simplified and unified lineage tracking with Global IDs
- More efficient job run time to significantly reduce job completion time for real world workloads
- Introducing Pachyderm deployment support using Helm Chart
- Updated python-pachyderm 7.0 release which supports Pachyderm 2.x

### Changes from Pachyderm 1.x behavior

- Empty directories are no longer supported
- Default upload behavior changes from append to overwrite
- Full paths in repo must be specified when uploading files
- Automatic file splitting is no longer supported
- `pachctl deploy` has been deprecated in favor of helm charts. See docs for deployment details
- Standby option replaced by autoscaling
- Writing multiple input datums to same output file is no longer supported and will result in an error 

Updated docs are at [docs.pachyderm.com](https://docs.pachyderm.com/latest/)

Try Pachyderm 2.0 on [hub.pachyderm.com](https://hub.pachyderm.com/landing?redirect=%2F)

Pachyderm 2.x is not backwards compatible with Pachyderm 1.x data formats. If you require assistance or have any questions, please contact [support@pachyderm.com](mailto:support@pachyderm.com)
