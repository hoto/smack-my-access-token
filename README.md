[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/smack-my-access-token/workflows/Build%20and%20test/badge.svg?branch=master)](https://github.com/hoto/smack-my-access-token/actions)
[![Release](https://img.shields.io/github/release/hoto/smack-my-access-token.svg?style=flat-square)](https://github.com/hoto/smack-my-access-token/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)

# Smack my login

Smackingtastic way of generating JWTs and HMAC signatures served over the web.  

Pure fantasy when using with postman in a pre-script:

```javascript
const getAccessTokenRequest = {
    url: 'http://localhost:5050/docusign',
    method: 'GET',
};

pm.sendRequest(getAccessTokenRequest, (err, res) => {
    console.log(err ? err : res.json());
    if (err === null) {
        var responseJson = res.json();
        pm.environment.set("currentAccessToken", responseJson.access_token);
    }
});
```

### Installation
    
Mac:

    brew install hoto/repo/smack-my-access-token

Mac or Linux:

    sudo curl -L \
      "https://github.com/hoto/smack-my-access-token/releases/download/1.0.0/smack-my-access-token_1.0.0_$(uname -s)_$(uname -m)" \
       -o /usr/local/bin/smack-my-access-token

    sudo chmod +x /usr/local/bin/smack-my-access-token
    
Or manually download binary from [releases](https://github.com/hoto/smack-my-access-token/releases).
    
### Run

    smack-my-access-token --help
    smack-my-access-token --version
    smack-my-access-token --port 5000

    curl localhost:5000

### Development

Build and test:

    go get github.com/hoto/smack-my-access-token
    
    make dependencies build test
    
Build binary:

     make build
    ./bin/smack-my-access-token

Run with arguments:

    make args="myargs" run

Install to global golang bin directory:

    make install
    smack-my-access-token
    
    