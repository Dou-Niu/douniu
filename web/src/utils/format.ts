import dayjs from 'dayjs'
const formatSeconds:(number)=>string=(time:number)=>{
  time = parseFloat(time.toFixed(2))
  let minutes:string | number = Math.floor(time / 60);
  let seconds:string | number  = Math.floor(time - minutes * 60);

  // 当秒数或者分数为0时，补0
  if(seconds < 10){
    seconds = "0" + seconds;
  }
  if(minutes < 10){
    minutes = "0" + minutes;
  }
  return minutes + ":" + seconds;
}

const formatTime = (timestamp:any)=>{
  return dayjs(timestamp).format('YYYY-MM-DD')
}

export {formatSeconds,formatTime}