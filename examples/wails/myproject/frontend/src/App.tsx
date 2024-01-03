import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { useGno } from "./api/use-gno";

function App() {
  const [resultText, setResultText] = useState([
    "Start GnoNative Kit on Wails to see result",
  ]);

  const gno = useGno();

  const onListAccounts = async () => {
    const res = await gno.listKeyInfo();
    if (res.length === 0) {
      setResultText([
        ...resultText,
        "No accounts found. Please create an account first.",
      ]);
    }
  };

  const onAction1 = async () => {
    const res = await gno.listKeyInfo();
    setResultText([...resultText, JSON.stringify(res, null, 2)]);

    // if (res.length > 0) {
    //   await gno.createAccount(
    //     "jefft0",
    //     "enable until hover project know foam join table educate room better scrub clever powder virus army pitch ranch fix try cupboard scatter dune fee",
    //     "password"
    //   );
    // }
    await gno.selectAccount("jefft0");
    await gno.setPassword("password");

    // const gasFee = "1000000ugnot";
    // const gasWanted = 2000000;
    // const args: Array<string> = ["3", "1", "1", "Hello all from Iuri"];
    // for await (const response of await gno.call(
    //   "gno.land/r/demo/boards",
    //   "CreateReply",
    //   args,
    //   gasFee,
    //   gasWanted
    // )) {
    //   setResultText(Buffer.from(response.result).toString());
    // }
  };

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div className="input-box"></div>
      <div id="result" className="result">
        {/* text area with multiple lines */}
        <textarea className="resultText"></textarea>
        {resultText.toString()}
      </div>
      <button className="btn" onClick={onListAccounts}>
        List Accounts
      </button>
      <button className="btn" onClick={onAction1}>
        Action 1
      </button>
    </div>
  );
}

export default App;
