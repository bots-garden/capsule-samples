## helloworld

- https://dev.to/satrobit/extending-envoy-with-webassembly-proxy-filters-1i96
- https://dev.to/mstryoda/extending-envoy-proxy-wasm-filter-with-golang-j8n

go get github.com/tetratelabs/proxy-wasm-go-sdk


./func-e run -c config.yaml



this example periodically logs the current time in nanoseconds

```
wasm log helloworld: OnPluginStart from Go!
wasm log helloworld: OnPluginStart from Go!
wasm log helloworld: OnTick on 1, it's 1601543078943978000
wasm log helloworld: OnTick on 1, it's 1601543078947916000
wasm log helloworld: OnTick on 1, it's 1601543078951979000
wasm log helloworld: OnTick on 1, it's 1601543079947462000
wasm log helloworld: OnTick on 1, it's 1601543079951503000
wasm log helloworld: OnTick on 1, it's 1601543079955484000
```
