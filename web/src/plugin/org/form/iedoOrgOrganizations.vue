<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="组织ID:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="组织名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入组织名称" />
       </el-form-item>
        <el-form-item label="组织编码:" prop="code">
          <el-input v-model="formData.code" :clearable="true"  placeholder="请输入组织编码" />
       </el-form-item>
        <el-form-item label="组织类型:" prop="type">
           <el-select v-model="formData.type" placeholder="请选择组织类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in orgtypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="上级组织:" prop="parentId">
          <el-tree-select
            v-model="formData.parentId"
            :data="orgTreeOptions"
            :props="{ label: 'name', value: 'id', children: 'children' }"
            placeholder="请选择上级组织"
            clearable
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="组织logo:" prop="logo">
          <el-input v-model="formData.logo" :clearable="true"  placeholder="请输入组织logo" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" :clearable="true">
            <el-option label="启用" value="1" />
            <el-option label="禁用" value="0" />
          </el-select>
       </el-form-item>
        <el-form-item label="联系人姓名:" prop="contactName">
          <el-input v-model="formData.contactName" :clearable="true"  placeholder="请输入联系人姓名" />
       </el-form-item>
        <el-form-item label="联系人电话:" prop="contactPhone">
          <el-input v-model="formData.contactPhone" :clearable="true"  placeholder="请输入联系人电话" />
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="组织描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入组织描述" />
       </el-form-item>
        <el-form-item label="排序:" prop="sort">
          <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createIedoOrgOrganizations,
  updateIedoOrgOrganizations,
  findIedoOrgOrganizations,
  getIedoOrgOrganizationsList
} from '@/plugin/org/api/iedoOrgOrganizations'

defineOptions({
    name: 'IedoOrgOrganizationsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const orgtypeOptions = ref([])
const orgTreeOptions = ref([])
const formData = ref({
  id: undefined,
  name: '',
  code: '',
  type: '',
  parentId: undefined,
  logo: '',
  status: '1',
  contactName: '',
  contactPhone: '',
  address: '',
  description: '',
  sort: 0,
})
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 列表转树型工具函数
function listToTree(list, parentId = 0) {
  return list
    .filter(item => item.parentId === parentId)
    .map(item => ({
      ...item,
      children: listToTree(list, item.id)
    }))
}

const loadOrgTree = async () => {
  const res = await getIedoOrgOrganizationsList({ page: 1, pageSize: 9999 })
  if (res.code === 0) {
    orgTreeOptions.value = listToTree(res.data.list)
  }
}

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findIedoOrgOrganizations({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    orgtypeOptions.value = await getDictFunc('orgtype')
}

init()
loadOrgTree()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createIedoOrgOrganizations(formData.value)
               break
             case 'update':
               res = await updateIedoOrgOrganizations(formData.value)
               break
             default:
               res = await createIedoOrgOrganizations(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
