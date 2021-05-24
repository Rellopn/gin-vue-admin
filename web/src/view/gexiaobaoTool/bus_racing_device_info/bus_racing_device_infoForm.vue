<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="设备号:">
                <el-input v-model="formData.imsi" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="topic:">
                <el-input v-model="formData.topic" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="消息体:">
                <el-input v-model="formData.payload" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="设备接收时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.messageData" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="接收到的时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.receiveDate" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="关卡id:"><el-input v-model.number="formData.racePassId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="比赛id:"><el-input v-model.number="formData.raceId" clearable placeholder="请输入"></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createRacingReveiveMessage,
    updateRacingReveiveMessage,
    findRacingReveiveMessage
} from "@/api/bus_racing_device_info";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "RacingReveiveMessage",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            imsi:"",
            topic:"",
            payload:"",
            messageData:new Date(),
            receiveDate:new Date(),
            racePassId:0,
            raceId:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createRacingReveiveMessage(this.formData);
          break;
        case "update":
          res = await updateRacingReveiveMessage(this.formData);
          break;
        default:
          res = await createRacingReveiveMessage(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findRacingReveiveMessage({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reracingReveiveMessage
       this.type = "update"
     }
    }else{
     this.type = "create"
   }
  
}
};
</script>

<style>
</style>