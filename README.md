# Introduction 
This project was created to have a easier way to provide our own extensions to the users of the IDE and have a better control of the versioning of each extension. That's why a custom extensions marketplace is a better solution for us to provide our extensions and keep control of some other (external) extension versions. And this also help us to allow users to only install the required extensions for a specific proyect. 

This is a fork of a open source repository [code-marketplace](https://github.com/coder/code-marketplace).

# Getting Started
You can read more information about the repository like the CLI commands of this project in the CODER_README.md within this project, these is the original README of the open source project. 

## Software dependencies
1. Install Docker CLI to build and publish Docker image.
2. To run it locally setup your environment to run Golang.

## Exposing the marketplace
When hosting the marketplace behind a reverse proxy set either the Forwarded header or both the X-Forwarded-Host and X-Forwarded-Proto headers. These headers are used to generate absolute URIs to extension assets in API responses. One way to test this is to make a query and check one of the URIs in the response:

```console
$ curl 'https://example.com/api/extensionquery' -H 'Accept: application/json;api-version=3.0-preview.1' --compressed -H 'Content-Type: application/json' --data-raw '{"filters":[{"criteria":[{"filterType":8,"value":"Microsoft.VisualStudio.Code"}],"pageSize":1}],"flags":439}' | jq .results[0].extensions[0].versions[0].assetUri
"https://example.com/assets/vscodevim/vim/1.24.1"
```

The marketplace does not support being hosted behind a base path; it must be proxied at the root of your domain.

# Build and Test

## Build
1. Build docker image: docker build ./ -t marketplace:0.0.0
2. Run docker container: docker run -p 80:80 marketplace:0.0.0  --extensions-dir /opt/source/extensions --address 0.0.0.0:80

# Contribute
1. Create a branch by following [our branching strategy](https://collaboration.artinsoft.com/tfs/Product/Bifrost/_wiki/wikis/Bifrost.wiki/880/Branching-Strategy).
2. Implement features according to our coding guidelines.
3. Submit a PR. Make sure that you specify [meaningful commit message](https://collaboration.artinsoft.com/tfs/Product/Bifrost/_wiki/wikis/Bifrost.wiki?wikiVersion=GBwikiMaster&pagePath=%2FBifrost%20project%20wiki%2FContribution%2FHow%20to%20Write%20Commit%20Message&pageId=881&_a=edit).
4. Make sure your PR has passed all quality gates (code review, static code analysis, unit tests, etc.).
5. Fix issues if any.
6. Merge PR.

If you want to learn more about creating good readme files then refer the following [guidelines](https://docs.microsoft.com/en-us/azure/devops/repos/git/create-a-readme?view=azure-devops). You can also seek inspiration from the below readme files:
- [ASP.NET Core](https://github.com/aspnet/Home)
- [Visual Studio Code](https://github.com/Microsoft/vscode)
- [Chakra Core](https://github.com/Microsoft/ChakraCore)

# How to upload extensions to the marketplace

1. Make sure that in the marketplace fileshare on the root folder exists a new-extensions folder.
2. If doesn't exists, create it. 
3. Upload the .vsix file to the new-extensions folder. 
4. Make a HTTP GET request to the endpoint [Market place url]/installextensions (could be done from a browser).
5. Watch the installation_logs.txt (inside the new-extensions folder) to keep track of the installation process.
