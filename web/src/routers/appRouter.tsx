import Home from "../pages/Home";
import ResourceBoard from "../pages/ResourceBoard";
import Signin from "../pages/Signin";

const appRouter = {
  "/resource": <ResourceBoard />,
  "/signin": <Signin />,
  "*": <Home />,
};

export default appRouter;
