import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
// 配置unocss
import UnoCSS from "unocss/vite";
import presetUno from "@unocss/preset-uno";
import presetAttributify from "@unocss/preset-attributify";
// 配置elementplus
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// 配置路径别名
import { resolve } from "path";
const pathResolve = (dir: string) => {
	return resolve(__dirname, ".", dir);
};
// https://vitejs.dev/config/
export default defineConfig({
	plugins: [vue(),
	UnoCSS({
		presets: [presetUno(), presetAttributify()]
	}),
	AutoImport({
		resolvers: [ElementPlusResolver()],
	}),
	Components({
		resolvers: [ElementPlusResolver()],
	})
	],
	resolve: {
		alias: {
			"@": pathResolve("src")
		}
	},
	server: {
		host: '127.0.0.1',
		port: 3000,
		proxy: {
			"/video": {
				target: "https://www.kecat.top/",
				ws: true,
				changeOrigin: true,
				rewrite: (path: any) => path.replace(/^\/video/, ""),
			},
			"/douniu/user": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
				// rewrite: (path: any) => path.replace(/^\/video/, ""),
			},
			"/douniu/video": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
				// rewrite: (path: any) => path.replace(/^\/video/, ""),
			},
			"/douniu/relation": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
				// rewrite: (path: any) => path.replace(/^\/video/, ""),
			},
			"/douniu/favourite": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
			},
			"/douniu/collection": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
			},
			"/douniu/comment": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
			},
			"/douniu/chat": {
				target: "http://43.139.195.17:8000",
				ws: true,
				changeOrigin: true,
			},
		},
	},
});
