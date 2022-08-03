import './wasm_exec'
import './wasmTypes.d.ts'

import React, { useEffect, useState } from 'react'

async function loadWasm(): Promise<void> {
  const goWasm = new window.Go();
  const result = await WebAssembly.instantiateStreaming(fetch('main.wasm'), goWasm.importObject);
  goWasm.run(result.instance)
}

export const LoadWasm: React.FC<React.PropsWithChildren<{}>> = (props) => {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    loadWasm().then(() => {
      setIsLoading(false);
    })
  }, []);

  if (isLoading) {
    return (
      <div>
        loading WebAssembly...
      </div>
    )
  } else {
    return <React.Fragment>{props.children}</React.Fragment>
  }
}
