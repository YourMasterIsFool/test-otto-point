import { Icon } from "@iconify/react/dist/iconify.js";
import React, { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";
import { ImageCanva } from "./ImageCanva";


interface IDropzoneInput {
    onChange:(imagePreview:string) => void
}
function DropzoneInput(props: IDropzoneInput) { 

    const {
        onChange
    } =  props
    // defining on drop change
    const onDrop = useCallback((acceptedFiles) => {
      const file = acceptedFiles[0];
      if (file) {
        const objectUrl = URL.createObjectURL(file);
        onChange(objectUrl)
      }
    }, []);
  const { getRootProps, getInputProps, acceptedFiles } = useDropzone({
    onDrop,
    accept: {
      "image/*": [],
    },
    multiple: false,
  });

  return (
    <div
      {...getRootProps()}
      className="bg-gray-100 w-full xl:py-6 lg:py-4 py-3 flex flex-col items-center"
    >

     
      <input {...getInputProps()} />
      <div className="flex flex-col w-full items-center space-y-2"> 
        <Icon icon={"bytesize:download"} fontSize={24} />
        <h1 className="xl:text-base text-sm">
            Browser File
        </h1>
        <p className="xl:text-base text-sm">
            Drag and Drop files
        </p>
      </div>
      {acceptedFiles.length > 0 && (
        <div>
          <ul>
            {acceptedFiles.map((file) => (
              <li key={file.path}>
                {file.path} - {file.size} bytes
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

export default DropzoneInput;
