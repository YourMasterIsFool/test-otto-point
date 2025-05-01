import { RefObject, useEffect, useRef } from "react";

interface IImageCanva<T> {
  src: string;
  dear: string;
  message: string;
  from: string;
  canvasRef: RefObject<T>
}

export function ImageCanva<T>({ src, dear, message, from, canvasRef }: IImageCanva<T>) {
//   const canvasRef = useRef<HTMLCanvasElement | null>(null);

  const draw = () => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const rect = canvas.getBoundingClientRect();

    // definisi canva heght dan width
    canvas.width = rect.width;
    canvas.height = rect.height;

    const ctx = canvas.getContext("2d");
    if (!ctx) return;

    // masukan image kedalam canva
    const image = new Image();
    image.src = src;

    image.onload = () => {
     // gambar canva
      ctx.drawImage(image, 0, 0, canvas.width, canvas.height);

      ctx.font = `${canvas.width * 0.03}px cursive`;

      ctx.fillText(`${dear}`, canvas.width * 0.45, canvas.height * 0.35);
      

      const lines = message.split("\n");
      lines.forEach((line, i) => {
        ctx.fillText(
          line,
          canvas.width * 0.3,
          canvas.height * 0.43 + i * (canvas.height * 1),
          canvas.width * 0.5
        );
      });

      ctx.fillText(`${from}`, canvas.width * 0.4, canvas.height * 0.6);
    };
  };

  useEffect(() => {

    // gambar kembali jika ada perubahan
    draw(); 

    const handleResize = () => {
      draw(); 
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, [src, dear, message, from]);

  return (
    <canvas
      ref={canvasRef}
      style={{
        width: "100%",
        height: "auto",
        maxWidth: "640px",
        display: "block",
        margin: "0 auto",
      }}
    />
  );
}
