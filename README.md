# GO versioning

Simple library for handling version numbers.

## Example

```go
// Start with 0.0.1
versionStruct = versioning.CreateVersion("0.0.1")
fmt.Printf("%s\n", versionStruct.GetVersionString())

// Make it 0.0.2
versionStruct.Patch()
fmt.Printf("%s\n", versionStruct.GetVersionString())

// Get TAG (for GIT Tagging)
fmt.Printf("%s\n", versionStruct.GetVersionTag())
```
