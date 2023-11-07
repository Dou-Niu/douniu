export const debounce = (fn: Function, delay: number) => {
    let timer: any = null;
    return function (...args: any[]) {
        clearTimeout(timer as any);
        timer = null;
        timer = setTimeout(() => {
            fn.apply(args)
            clearTimeout(timer)
            timer = null;
        }, delay);
    }
}