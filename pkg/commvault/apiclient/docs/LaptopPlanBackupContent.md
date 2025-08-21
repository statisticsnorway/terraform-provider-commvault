# LaptopPlanBackupContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsIncludedPaths** | **[]string** | Paths to include for Windows | [optional] [default to ["Desktop","Documents","MigrationAssistant"]]
**WindowsExcludedPaths** | **[]string** | Paths to exclude for Windows | [optional] [default to ["<WKF,AppData>","<WKF,LocalAppData>","Disk Images","Executable","Temporary Files (Windows)","C:\\Program Files","C:\\Program Files (x86)","C:\\Windows"]]
**MacIncludedPaths** | **[]string** | Paths to include for Mac | [optional] [default to ["Desktop","Documents","MigrationAssistant"]]
**MacExcludedPaths** | **[]string** | Paths to exclude for Mac | [optional] [default to ["Disk Images","Executable","Temporary Files (Mac)","/Library","<WKF,Library>"]]
**UnixIncludedPaths** | **[]string** | Paths to include for UNIX | [optional] [default to ["Desktop","Documents"]]
**UnixExcludedPaths** | **[]string** | Paths to exclude for UNIX | [optional] [default to ["Disk Images","Executable","Temporary Files (Linux)"]]
**FileSystemQuota** | **int32** | Maximum number of gigabytes that you can store in the File System. Giving value as -1 means infinite file system quota. | [optional] [default to -1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

