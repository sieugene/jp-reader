import { OcrData } from "@/shared/services/mokuro";
import { FC } from "react";

type Props = {
  imageSrc: string;
  data: OcrData;
  scaleX?: number;
  scaleY?: number;
};
export const Reader: FC<Props> = ({
  imageSrc,
  scaleX = 1,
  scaleY = 1,
  data,
}) => {
  const { img_width, img_height, blocks } = data;

  const textBlocks = data.blocks.map((block) => block.lines.join("\n"));

  return (
    <div>
      <div
        style={{
          position: "relative",
          width: `${img_width * scaleX}px`,
          height: `${img_height * scaleY}px`,
          backgroundImage: `url(${imageSrc})`,
          backgroundSize: "contain",
          backgroundRepeat: "no-repeat",
        }}
      >
        {blocks.map((block, index) => (
          <div
            key={index}
            style={{
              position: "absolute",
              left: `${block.box[0] * scaleX}px`,
              top: `${block.box[1] * scaleY}px`,
              width: `${(block.box[2] - block.box[0]) * scaleX}px`,
              height: `max-content`,
              writingMode: block.vertical ? "vertical-rl" : "horizontal-tb",
              fontSize: `${block.font_size * Math.min(scaleX, scaleY)}px`,
              color: "black",
              textAlign: "justify",
              background: "white",
            }}
          >
            {block.lines.map((line, lineIndex) => (
              <div key={lineIndex}>{line}</div>
            ))}
          </div>
        ))}
      </div>
      <div className="text-blocks">
        {textBlocks.map((block, index) => (
          <div key={index} style={{ border: "1px solid black" }}>
            <p>{block}</p>
          </div>
        ))}
      </div>
    </div>
  );
};
