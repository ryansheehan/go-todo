import { resolve } from 'path';
import { defineConfig } from 'vite';

export default defineConfig({
    cacheDir: "tmp/web",
    publicDir: "assets",
    build: {
        lib: {
            entry: [resolve(__dirname, "web/main.ts")],
            formats: ["es"],
            name: "[name]",
            fileName: "[name]",
        },
        outDir: "static",
        emptyOutDir: true,
        watch: {            
            include: ["web/**"]
        },
    },
});
