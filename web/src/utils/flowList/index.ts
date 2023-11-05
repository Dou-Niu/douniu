import { Ref, ref } from 'vue'

// 获得最小高度下标
function getMin(arr: number[]) {
    let minAns = 0;
    for (let i = 1; i < arr.length; i++) {
        if (arr[i] < arr[minAns]) {
            minAns = i;
        }
    }
    return minAns;
}

export class flowList {
    curSpace: number
    curRenderListInfo: Ref<any[]>
    itemWidth: number
    public heights: number[]
    colCount: number
    rowSpace: number
    curContainerWidth: number
    curIndex: number
    maxHeight: Ref<number>
    constructor(list: any[], itemWidth: number, rowSpace?: number) {
        this.curSpace = 0;
        this.curRenderListInfo = ref(list);
        this.itemWidth = itemWidth / 100;
        this.colCount = 0;
        this.heights = Array<number>(this.colCount).fill(0);
        // 默认上下列表项之间是5px空隙
        this.rowSpace = rowSpace || 5;
        this.curContainerWidth = 0;
        this.curIndex = 0;
        this.maxHeight = ref(0);
    }
    // 计算各项之间空隙和列数
    calcSpaceCols(containerWidth: number, itemWidth: number = this.itemWidth) {
        this.curContainerWidth = containerWidth;
        this.itemWidth = itemWidth;
        this.colCount = Math.floor(containerWidth / (itemWidth * containerWidth));
        this.curSpace = (containerWidth - (itemWidth * containerWidth) * this.colCount) / (this.colCount - 1);
        this.heights = Array<number>(this.colCount).fill(0);
    }

    /**
     * 列表项初始化时，将列表项赋予top和left定位
     * @param item 列表项，通过它找到对应列表项下标
     * @param itemDom 列表项DOM元素
     */
    setPosition(item: any, itemDom: HTMLElement) {
        return new Promise((resolve) => {
            itemDom.style.transition = 'none';
            itemDom.style.transform = `translate(calc(${this.curContainerWidth / 2}px - 50%),${this.maxHeight.value}px)`;
            setTimeout(() => {
                itemDom.style.transition = 'transform .8s';
                let ind = this.curRenderListInfo.value.indexOf(item);
                let height = itemDom.offsetHeight;
                const pos = getMin(this.heights);
                this.curRenderListInfo.value[ind].top = this.heights[pos] === 0 ? 0 : this.heights[pos] + this.rowSpace;
                this.heights[pos] += this.heights[pos] === 0 ? height : height + this.rowSpace;
                this.curRenderListInfo.value[ind].left = pos === 0 ? 0 : pos * (this.itemWidth * this.curContainerWidth + this.rowSpace);
                this.curRenderListInfo.value[ind].flow_index = this.curIndex;
                this.curIndex++;
                this.curRenderListInfo.value[ind].dom = itemDom;
                this.maxHeight.value = Math.max(...this.heights);
                resolve('success');
            }, 10);
        })
    }
    updatePosition(item: any, itemDom: HTMLElement) {
        let ind = this.curRenderListInfo.value.indexOf(item);
        let height = itemDom.offsetHeight;
        const pos = getMin(this.heights);
        this.curRenderListInfo.value[ind].top = this.heights[pos] === 0 ? 0 : this.heights[pos] + this.rowSpace;
        this.heights[pos] += this.heights[pos] === 0 ? height : height + this.rowSpace;
        this.curRenderListInfo.value[ind].left = pos === 0 ? 0 : pos * (this.itemWidth * this.curContainerWidth + this.rowSpace);
        this.curRenderListInfo.value[ind].flow_index = this.curIndex;
        this.curIndex++;
        this.curRenderListInfo.value[ind].dom = itemDom;
        this.maxHeight.value = Math.max(...this.heights);
    }
    /**
     * 窗口尺寸变化的时候，更新布局
     * @param updatedContainerWidth 更新后窗口容器宽度
     */
    updateOnResize(updatedContainerWidth: number, itemWidth: number = this.itemWidth) {
        this.calcSpaceCols(updatedContainerWidth, itemWidth);
        this.curRenderListInfo.value.sort((a: any, b: any) => {
            return a.flow_index - b.flow_index;
        })
        this.curRenderListInfo.value.forEach(item => {
            this.updatePosition(item, item.dom);
        })
    }
}


