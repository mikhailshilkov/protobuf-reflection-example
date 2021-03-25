# Protobuf Reflection Example

A Go example that dynamically loads proto files and outputs metadata, including comments.

Program output:

```
=====
A service that application uses to manipulate triggers and functions.
=====

ListFunctions:
 Returns a list of functions that belong to the requested project.

GetFunction:
 Returns a function with the given name from the requested project.

CreateFunction:
 Creates a new function. If a function with the given name already exists in
 the specified project, the long running operation will return
 `ALREADY_EXISTS` error.

UpdateFunction:
 Updates existing function.

DeleteFunction:
 Deletes a function with the given name from the specified project. If the
 given function is used by some trigger, the trigger will be updated to
 remove this function.

CallFunction:
 Synchronously invokes a deployed Cloud Function. To be used for testing
 purposes as very limited traffic is allowed. For more information on
 the actual limits, refer to
 [Rate Limits](https://cloud.google.com/functions/quotas#rate_limits).

GenerateUploadUrl:
 Returns a signed URL for uploading a function source code.
 For more information about the signed URL usage see:
 https://cloud.google.com/storage/docs/access-control/signed-urls.
 Once the function source code upload is complete, the used signed
 URL should be provided in CreateFunction or UpdateFunction request
 as a reference to the function source code.

 When uploading source code to the generated signed URL, please follow
 these restrictions:

 * Source file type should be a zip file.
 * Source file size should not exceed 100MB limit.
 * No credentials should be attached - the signed URLs provide access to the
   target bucket using internal service identity; if credentials were
   attached, the identity from the credentials would be used, but that
   identity does not have permissions to upload files to the URL.

 When making a HTTP PUT request, these two headers need to be specified:

 * `content-type: application/zip`
 * `x-goog-content-length-range: 0,104857600`

 And this header SHOULD NOT be specified:

 * `Authorization: Bearer YOUR_TOKEN`

GenerateDownloadUrl:
 Returns a signed URL for downloading deployed function source code.
 The URL is only valid for a limited period and should be used within
 minutes after generation.
 For more information about the signed URL usage see:
 https://cloud.google.com/storage/docs/access-control/signed-urls

SetIamPolicy:
 Sets the IAM access control policy on the specified function.
 Replaces any existing policy.

GetIamPolicy:
 Gets the IAM access control policy for a function.
 Returns an empty policy if the function exists and does not have a policy
 set.

TestIamPermissions:
 Tests the specified permissions against the IAM access control policy
 for a function.
 If the function does not exist, this will return an empty set of
 permissions, not a NOT_FOUND error.
```
