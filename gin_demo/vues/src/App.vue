<template>
  <div>
    <h1 align="center" class="title">支持数据价值交换的区块链分链管理软件</h1>
    <el-row>
      <el-col :span="6">
        <el-card class="in-card">
          <div slot="header" class="clearfix">
            <span>事务信息输入</span>
          </div>
          <el-form ref="inform" :model="inform" label-width="150px" >
            <el-form-item label="数据价值ID" style="width:85%">
              <el-input v-model="inform.id"></el-input>
            </el-form-item>
            <el-form-item label="数据价值Price" style="width:85%">
              <el-input v-model="inform.price"></el-input>
            </el-form-item>
            <el-form-item label="数据价值Label" style="width:85%">
              <el-input v-model="inform.label"></el-input>
            </el-form-item>
            <el-form-item label="分链方式">
              <el-radio-group v-model="inform.chainID">
                <el-radio label="ethereum"></el-radio>
                <el-radio label="hyperledger fabric"></el-radio>
                <el-radio label="auto"></el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="onSubmit">提交</el-button>
              <el-button @click="onCancel">取消</el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="6"  >
        <el-card class="out-card" >
          <div slot="header" class="clearfix">
            <span>事务结果</span>
          </div>
          <div class="result-show">事务状态： {{outform.status}}</div>
          <div class="result-show">txID： {{outform.txID}}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>

</template>

<script>
import axios from 'axios'
export default {
  name: 'App',
  methods: {
    async onSubmit () {
      this.fullscreenLoading = true
      const loading = this.$loading({
        lock: true,
        text: '正在获取结果',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      const res = await axios.get('http://127.0.0.1:8000/v1/onchain', {
        params: {
          id: this.inform.id,
          price: this.inform.price,
          label: this.inform.label,
          chainID: this.inform.chainID
        }
      })
      this.outform.status = res.data.status
      this.outform.txID = res.data.txID
      this.outform.blockID = res.data.blockID
      this.fullscreenLoading = false
      loading.close()
    },
    onCancel () {
      console.log('clearAll')
    }
    // onSubmit () {
    //   console.log('submit')
    // }
  },
  data () {
    return {
      fullscreenLoading: false,
      txID: '!@##',
      blockID: '##@!',
      inform: {
        id: '',
        price: '',
        label: '',
        chainID: ''
      },
      outform: {
        status: '',
        txID: '',
        blockID: ''
      }
    }
  }
}
</script>

<style>
.in-card {
  width: 600px;
  margin-top: 50px;
  margin-left: 300px;
}
.out-card {
  width: 600px;
  margin-top: 50px;
  margin-left: 500px;
}
.title {
  margin-top: 50px;
}
</style>
