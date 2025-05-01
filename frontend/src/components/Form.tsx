import { Icon } from "@iconify/react";
import DropzoneInput from "./DropzoneInput";
import { useForm } from "react-hook-form";
import { ImageCanva } from "./ImageCanva";
import { useRef, useState } from "react";


interface IFormInput {
    file:string,
    dear:string,
    message: string,
    from:string
}
export function Form() {

const canvasRef = useRef(null)
  const { register,
    formState: {
        errors,
    },
    handleSubmit,setValue, watch } = useForm<IFormInput>();

    // download ketika telah semuanya validation lewad
  function onSubmit(value) {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const link = document.createElement("a");
    link.download = "card.png";
    link.href = canvas.toDataURL("image/png");
    link.click();
  }



 return (
   <form className="w-full xl:p-6 p-4 bg-white" onSubmit={handleSubmit(onSubmit)}>
     <h1 className="font-bold text-center xl:text-2xl text-xl">Gift Card</h1>
     <div className="flex border-t-1  border-b-1 border-gray-200 lg:my-8 py-4 lg:py-8 pt-6 flex-col w-full items-start space-y-4">
       {watch("file") ? (
         <ImageCanva
         canvasRef={canvasRef}
        //    onChange={(value) => onChangeImageCanva(value)}
           dear={watch("dear")}
           message={watch("message")}
           from={watch("from")}
           src={watch("file")}
         />
       ) : null}
       <div className="relative xl:mt-12 mt-6 w-full">
         <p>File Upload</p>

         <DropzoneInput
           onChange={(imagePreviewUrl) => setValue("file", imagePreviewUrl)}
         />
       </div>

       <div className="form-group">
         <label>Dear</label>
         <input
           {...register("dear", {
             required: true,
           })}
           className="form-input"
         />

         {errors.dear?.type === "required" && (
           <p className="error-message">Dear is required</p>
         )}
       </div>

       <div className="form-group">
         <label>Message</label>
         <input
           {...register("message", {
             required: true,
           })}
           className="form-input"
         />
         {errors.message?.type === "required" && (
           <p className="error-message">Message is required</p>
         )}
       </div>

       <div className="form-group">
         <label>From</label>
         <input
           {...register("from", {
             required: true,
           })}
           className="form-input"
         />
         {errors.from?.type === "required" && (
           <p className="error-message">From is required</p>
         )}
       </div>
     </div>
     <div className="flex justify-center ">
       <button className="button button-primary">Download</button>
     </div>
   </form>
 );
}
