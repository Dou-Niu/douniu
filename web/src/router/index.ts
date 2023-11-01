import { createRouter, createWebHashHistory} from "vue-router";
import {HomeRoutes,ErrRoutes} from "./routes";

const router = createRouter({
	history: createWebHashHistory(),
	routes: [...HomeRoutes,...ErrRoutes]
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

// router.afterEach(async (to, from, failure) => {});

router.onError((error) => {
	console.log(error, "路由错误");
});


export default router;
