import { useState } from 'react'

import { Form } from './components/Form'
function App() {
  return (
    <>
      <div className="md:flex w-full h-full xl:pt-8 pt-4  md:justify-center">
        <div className="fixed bg-indigo-50 top-0 left-0 z-0 h-screen w-screen"></div>
        <div className="md:w-1/2 relative w-full px-6">
          <Form />
        </div>
      </div>
    </>
  );
}

export default App
