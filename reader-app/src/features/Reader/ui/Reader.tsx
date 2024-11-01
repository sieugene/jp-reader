
import { TransformedProject } from '@/shared/services/api';
import { FC } from "react";

type Props = {
  scaleX?: number;
  scaleY?: number;
} & TransformedProject["data"][0];

export const Reader: FC<Props> = ({
  scaleX = 1,
  scaleY = 1,
  image,
  ocrData,
}) => {
  const { img_width, img_height, blocks } = ocrData;

  const textBlocks = ocrData.blocks.map((block) => block.lines.join("\n"));

  return (
    <div className="reader-container p-4 bg-base-100 rounded-lg shadow-lg">
      <div
        style={{
          position: "relative",
          width: `${img_width * scaleX}px`,
          height: `${img_height * scaleY}px`,
          backgroundImage: `url(${image})`,
          backgroundSize: "contain",
          backgroundRepeat: "no-repeat",
          border: "1px solid #e0e0e0",
          borderRadius: "8px",
          overflow: "hidden",
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
              background: "rgba(255, 255, 255, 0.7)",
              borderRadius: "4px",
              boxShadow: "0 2px 4px rgba(0, 0, 0, 0.2)",
            }}
          >
            {block.lines.map((line, lineIndex) => (
              <div key={lineIndex}>{line}</div>
            ))}
          </div>
        ))}
      </div>

      <div className="text-blocks mt-4">
        {textBlocks.map((block, index) => (
          <div
            key={index}
            className="border border-base-300 rounded-lg p-2 mb-2 bg-base-200"
          >
            <p className="text-gray-800">{block}</p>
          </div>
        ))}
      </div>
    </div>
  );
};
