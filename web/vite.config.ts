import { defineConfig } from "@solidjs/start/config";
import devtoolsPlugin from "solid-devtools/vite";

export default defineConfig({
    plugins: [
        devtoolsPlugin({
            autoname: true,
        }),
    ]
});
