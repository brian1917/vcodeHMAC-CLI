# Veracode HMAC CLI

## Description
CLI tool to generate an authorization header for Veracode APIs using API ID and Key.

## Third-party Packages
github.com/brian1917/vcodeHMAC (https://godoc.org/github.com/brian1917/vcodeHMAC)

## Parameters
1. `-credsFile`: path to the file with Veracode credentials (see details below)
2. `-url`: The URL for the request (e.g., `https://analysiscenter.veracode.com/api/5.0/detailedreport.do?build_id=123456`)
3. `-method`: HTTP Method (e.g., GET, POST, etc.)

## Credentials File
Must be structured like the following:
```
[DEFAULT]
veracode_api_key_id = ID HERE
veracode_api_key_secret = SECRET HERE
```

## Executables
I've added the executables for Mac (vcodeHMAC-CLI) and Windows (vcodeHMAC-CLI.exe).
Building from source is preferred, but I'll try to keep these up-to-date for those that don't have Go installed.
* For Mac, download the executable, set it to be an executable: `chmod +x vcodeHMAC-CLI`
