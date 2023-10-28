import { createApp } from "vue";
import "uno.css";
import App from "./App.vue";
import router from "./router";
// 引入elementPlus图标
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

const app = createApp(App);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
	app.component(key, component);
}
// 引入pinia
import { createPinia } from "pinia";
const pinia = createPinia();

app.use(pinia);
app.use(router);
app.mount("#app");
