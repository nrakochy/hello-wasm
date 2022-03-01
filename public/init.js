'use strict';

const WASM_URL = 'threader.wasm';

(function init() {
    const go = new Go();

    if ('instantiateStreaming' in WebAssembly) {
        WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
            console.error("FOUND IT ", obj.instance)
            const wasm = obj.instance;
            window.wasm = wasm;
            go.run(wasm);
        })
    }
})();