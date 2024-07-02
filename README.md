# ZAP GO API

The Go implementation to access the [ZAP API](https://www.zaproxy.org/docs/api/). For more information
about ZAP consult the (main) [ZAP project](https://github.com/zaproxy/zaproxy/).

## How to Obtain

The latest released version can be downloaded using:

    go get github.com/zaproxy/zap-api-go

## Getting Help

For help using ZAP API refer to:
  * [Examples](https://github.com/zaproxy/zap-api-go/tree/master/example) - collection of examples using the library;
  * [API Documentation](https://www.zaproxy.org/docs/api/)
  * [ZAP User Group](https://groups.google.com/group/zaproxy-users) - for asking questions;

## Updating the Generated Files

Most of the API code is generated from the ZAP java source code.

To regenerate the API code you will need the repos [zaproxy](https://github.com/zaproxy/zaproxy) and [zap-extensions](https://github.com/zaproxy/zap-extensions) checked out at the same level as this one.

Cloning the Repositories:
```
git clone --recursive -j8 https://github.com/zaproxy/zaproxy.git
git clone --recursive -j8 https://github.com/zaproxy/zap-extensions.git
```

Typically, you should generate the core API calls from the latest release tag. For example:

```
cd zaproxy
git fetch upstream -t
git checkout tags/v2.15.0
./gradlew generateGoApiEndpoints
cd ..
```

The add-on APIs can be generated from the zap-extensions main branch:

```
cd zap-extensions
./gradlew generateGoZapApiClientFiles --continue
cd ..
```

Finally, run the command to update the `interface.go`:

```
/bin/bash zap-api-go/zap/generate_interface.sh
```

The above commands will update the files in `zap-api-go/zap`.

If any new files are created then they should be manually added to `zap-api-go/zap/interface.go` as per the existing files.