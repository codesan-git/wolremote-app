import { useState } from "react";
import { Command } from "@tauri-apps/plugin-shell";

// --- KONFIGURASI MAPPING DI SINI ---
const CONFIG_BASE_PATH = "/Users/pc/Developers/misc/wolremote-v2/wolremote-gui/configs";

const CONFIG_LIST = [
  { label: "PC 1 (Server)", filename: "config_1.toml" },
  { label: "PC 2", filename: "config_2.toml" },
];

type ActionType = "on" | "off" | "status" | "connect" | "disconnect" | "run" | "check" | "stop";

function App() {
  const [output, setOutput] = useState<string>("Ready...");
  const [isLoading, setIsLoading] = useState(false);

  // State untuk menyimpan file yang dipilih dari dropdown
  const [selectedFile, setSelectedFile] = useState(CONFIG_LIST[0].filename);

  const runWolRemote = async (action: ActionType) => {
    setIsLoading(true);

    // Gabungkan Base Path + Nama File
    // Hasilnya: /Users/pc/.../configs/config_1.toml
    // Kita tambahkan "/" manual untuk jaga-jaga
    const fullPath = `${CONFIG_BASE_PATH}/${selectedFile}`.replace('//', '/');

    setOutput(`Target: ${selectedFile}\nExecuting: wolremote ${action}...`);

    try {
      const command = Command.sidecar("binaries/wolremote", [
        action,
        "--remote",
        fullPath,
      ]);

      const result = await command.execute();

      if (result.code === 0) {
        setOutput(result.stdout || "Success (No Output)");
      } else {
        setOutput(`Error (Code ${result.code}):\n${result.stderr}`);
      }
    } catch (err) {
      setOutput(`System Error: ${String(err)}`);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen p-8 flex flex-col items-center justify-center gap-6 bg-slate-900 text-slate-100 font-sans">
      <h1 className="text-3xl font-bold text-teal-400 tracking-wide">
        WOL CONTROL CENTER
      </h1>

      {/* Dropdown Selection */}
      <div className="w-full max-w-md bg-slate-800 p-4 rounded-lg border border-slate-700 shadow-lg">
        <label className="text-xs text-slate-400 uppercase font-bold block mb-2">
          Select Target Machine
        </label>

        <div className="relative">
          <select
            value={selectedFile}
            onChange={(e) => setSelectedFile(e.target.value)}
            className="w-full appearance-none bg-slate-900 border border-slate-600 text-white py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:border-teal-500 cursor-pointer hover:bg-slate-950 transition"
          >
            {CONFIG_LIST.map((config) => (
              <option key={config.filename} value={config.filename}>
                {config.label}
              </option>
            ))}
          </select>
          {/* Panah Dropdown Custom */}
          {/* <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-slate-400">
            <svg className="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
              <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
            </svg>
          </div> */}
        </div>

        <div className="mt-2 text-[10px] text-slate-500 font-mono truncate">
          Path: {CONFIG_BASE_PATH}/{selectedFile}
        </div>
      </div>

      {/* Control Grid */}
      <div className="grid grid-cols-2 gap-4 w-full max-w-md">
        <ActionButton action="on" onClick={() => runWolRemote("on")} color="bg-green-600 hover:bg-green-500" />
        <ActionButton action="off" onClick={() => runWolRemote("off")} color="bg-red-600 hover:bg-red-500" />

        <div className="col-span-2 h-px bg-slate-700 my-2"></div>

        <ActionButton action="status" onClick={() => runWolRemote("status")} color="bg-blue-600 hover:bg-blue-500" />
        <ActionButton action="check" onClick={() => runWolRemote("check")} color="bg-cyan-600 hover:bg-cyan-500" />

        <ActionButton action="connect" onClick={() => runWolRemote("connect")} color="bg-indigo-600 hover:bg-indigo-500" />
        <ActionButton action="disconnect" onClick={() => runWolRemote("disconnect")} color="bg-purple-600 hover:bg-purple-500" />

        <ActionButton action="run" onClick={() => runWolRemote("run")} color="bg-orange-600 hover:bg-orange-500" />
        <ActionButton action="stop" onClick={() => runWolRemote("stop")} color="bg-rose-600 hover:bg-rose-500" />
      </div>

      {/* Terminal Output */}
      <div className="w-full max-w-md bg-black rounded-lg border border-slate-800 p-4 font-mono text-xs h-48 overflow-auto shadow-xl">
        <div className="flex justify-between mb-2">
          <span className="text-slate-500">Terminal Output</span>
          {isLoading && <span className="text-yellow-400 animate-pulse">Processing...</span>}
        </div>
        <pre className="whitespace-pre-wrap text-slate-300">
          {output}
        </pre>
      </div>
    </div>
  );
}

function ActionButton({ action, onClick, color }: { action: string, onClick: () => void, color: string }) {
  return (
    <button
      onClick={onClick}
      className={`py-3 px-4 rounded font-semibold text-white uppercase tracking-wider shadow-md transition active:scale-95 active:shadow-inner ${color}`}
    >
      {action}
    </button>
  );
}

export default App;