/** @type {import('vite').UserConfig} */
export default {
    cacheDir: "tmp/web",
    build: {
        outDir: "public",
        rollupOptions: {
            input: {
                main: './web/main.ts'
            }
        }
    }
}
