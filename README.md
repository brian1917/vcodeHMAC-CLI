# Veracode HMAC CLI

## Description
CLI tool to generate an authorization header for Veracode APIs using API ID and Key. Given an HTTP method and URL, and the location of your Veracode API credentials file, you will get the value of an Authorization header printed out for piping into curl, httpie, or other scripting uses.

## Third-party Packages
github.com/brian1917/vcodeHMAC (https://godoc.org/github.com/brian1917/vcodeHMAC)

## Parameters
1. `-credsFile`: path to the file with Veracode credentials (see details below)
2. `-url`: The URL for the request (e.g., `https://analysiscenter.veracode.com/api/5.0/detailedreport.do?build_id=123456`)
3. `-method`: HTTP Method (e.g., GET, POST, etc.)

## Executables
Executables for Windows, Mac, and Linux are available in the releases section of this repository: https://github.com/brian1917/vcodeHMAC-CLI/releases.

## Credentials File
Must be structured like the following:
```
[DEFAULT]
veracode_api_key_id = ID HERE
veracode_api_key_secret = SECRET HERE
```