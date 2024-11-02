import { Box } from "@mui/material"
import Form from "./components/form"
import Header from "./components/header"

function App() {

  return (
    <>
      <Box position={"fixed"}>
        <Header />
      </Box>
      <Box display={"flex"} justifyContent={"center"} alignItems={"center"} height={"100vh"}>
        <Form />
    </Box>
    </>
  )
}

export default App
