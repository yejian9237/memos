import { useEffect } from "react";
import "../less/resource-board.less";

interface Props {}

const ResourceBoard: React.FC<Props> = () => {
  useEffect(() => {
    // do nth
  }, []);

  return (
    <div className="page-wrapper resource-board-wrapper">
      <p>This is resoure page</p>
    </div>
  );
};

export default ResourceBoard;
