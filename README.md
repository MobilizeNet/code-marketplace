

# Contribute

## How to upload extensions to the marketplace

1. Make sure that in the marketplace fileshare on the root folder exists a new-extensions folder.
2. If doesn't exists, create it. 
3. Upload the .vsix file to the new-extensions folder. 
4. Make a HTTP GET request to the endpoint [Market place url]/installextensions (could be done from a browser).
5. Watch the installation_logs.txt (inside the new-extensions folder) to keep track of the installation process.
