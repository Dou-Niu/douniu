import { createRouter, createWebHashHistory} from "vue-router";
import {HomeRoutes,ErrRoutes,ChannelRoutes} from "./routes";

const router = createRouter({
	history: createWebHashHistory(),
	routes: [...HomeRoutes,...ErrRoutes,...ChannelRoutes]
});

// 路由守卫

router.beforeEach(async (to, _, next) => {
	if (!to.name) {
		//判断有没有路由
		next({ name: "404" });
	} else {
		next();
	}
});

router.onError((error) => {
	console.log(error, "路由错误");
});


export default router;
