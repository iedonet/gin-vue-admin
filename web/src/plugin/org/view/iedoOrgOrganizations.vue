<template>
  <div class="org-layout">
    <el-row :gutter="16" class="org-main-row">
      <!-- 左侧组织树 -->
      <el-col :span="5" class="org-tree-col">
        <el-card shadow="hover" class="org-tree-card">
          <div class="org-tree-header">
            <span>组织架构</span>
            <el-button type="primary" size="small" @click="addRootOrg">新增根组织</el-button>
          </div>
          <el-tree
            :data="treeData"
            :props="TREE_PROPS"
            node-key="id"
            highlight-current
            @node-click="handleNodeClick"
            :expand-on-click-node="false"
            class="org-tree"
            :default-expanded-keys="defaultExpandedKeys"
          >
            <template #default="{ node, data }">
              <div class="custom-tree-node">
                <span class="node-label">{{ data.name }}</span>
                <el-tag size="small" :type="getOrgTypeTagType(data.type)" style="margin-left: 8px;">
                  {{ orgtypeOptions.find(opt => opt.value === data.type)?.label || '-' }}
                </el-tag>
                <div class="node-actions">
                  <el-icon class="action-icon" @click.stop="addChildOrg(data)"><CirclePlus /></el-icon>
                  <el-icon class="action-icon" @click.stop="editOrg(data)"><Edit /></el-icon>
                  <el-icon class="action-icon delete" @click.stop="deleteOrg(data)" v-if="!data.children?.length"><Delete /></el-icon>
                </div>
              </div>
            </template>
          </el-tree>
        </el-card>
      </el-col>
      <!-- 中间组织详情 -->
      <el-col :span="7" class="org-detail-col">
        <el-card shadow="hover" class="org-detail-card">
          <template v-if="currentOrg">
            <div class="org-detail-header">
              <span>组织详情</span>
              <div class="org-header-actions">
                <el-button type="primary" size="small" @click="addChildOrg(currentOrg)">新增下级</el-button>
                <el-button type="primary" size="small" @click="editOrg(currentOrg)">编辑</el-button>
              </div>
            </div>
            <el-descriptions :column="1" border class="org-desc">
              <el-descriptions-item label="名称">{{ currentOrg.name }}</el-descriptions-item>
              <el-descriptions-item label="编码">{{ currentOrg.code }}</el-descriptions-item>
              <el-descriptions-item label="类型">{{ filterDict(currentOrg.type, orgtypeOptions) }}</el-descriptions-item>
              <el-descriptions-item label="Logo">
                <el-image v-if="currentOrg.logo" :src="currentOrg.logo" style="width:60px;height:60px;" />
                <span v-else>无Logo</span>
              </el-descriptions-item>
              <el-descriptions-item label="状态">{{ currentOrg.status }}</el-descriptions-item>
              <el-descriptions-item label="联系人">{{ currentOrg.contactName }}</el-descriptions-item>
              <el-descriptions-item label="电话">{{ currentOrg.contactPhone }}</el-descriptions-item>
              <el-descriptions-item label="地址">{{ currentOrg.address }}</el-descriptions-item>
              <el-descriptions-item label="描述">{{ currentOrg.description }}</el-descriptions-item>
              <el-descriptions-item label="排序">{{ currentOrg.sort }}</el-descriptions-item>
            </el-descriptions>
          </template>
          <template v-else>
            <el-empty description="请在左侧选择一个组织节点" />
          </template>
        </el-card>
      </el-col>
      <!-- 右侧成员管理（可留空或mock） -->
      <el-col :span="12" class="org-member-col">
        <el-card shadow="hover" class="org-member-card">
          <div class="org-member-header">
            <span>成员管理</span>
            <el-button type="primary" size="small">添加成员</el-button>
          </div>
          <el-empty description="请完善成员管理功能" />
        </el-card>
      </el-col>
    </el-row>
    <!-- 组织编辑/新增弹窗 -->
    <el-dialog v-model="dialogFormVisible" :title="type === 'create' ? '新增组织' : '编辑组织'" width="600px">
      <el-form ref="elFormRef" :model="formData" :rules="RULES" label-width="80px">
        <el-form-item label="上级组织" prop="parentId" v-if="dialogFormVisible">
          <el-tree-select
            
          v-model="formData.parentId"
            :data="treeData"
            :props="TREE_PROPS"
            clearable
            style="width: 100%"
            placeholder="请选择上级组织"
          />
        </el-form-item>
        <el-form-item label="组织名称" prop="name">
          <el-input v-model="formData.name" placeholder="请输入组织名称" />
        </el-form-item>
        <el-form-item label="组织类型" prop="type">
          <el-select v-model="formData.type" placeholder="请选择组织类型" style="width: 100%">
            <el-option v-for="item in orgtypeOptions" :key="item.value" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="Logo" prop="logo">
          <div style="display: flex; align-items: center; gap: 10px;">
            <el-input v-model="formData.logo" placeholder="Logo URL" style="width: 300px;" />
            <el-upload
              class="logo-uploader"
              :action="uploadUrl"
              :show-file-list="false"
              :on-success="handleLogoSuccess"
              :before-upload="beforeLogoUpload"
              :data="{ orgId: formData.id }"
            >
              <el-button size="small" type="primary">上传Logo</el-button>
            </el-upload>
          </div>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="formData.status" placeholder="请选择状态" style="width: 100%">
            <el-option label="启用" value="1" />
            <el-option label="禁用" value="0" />
            <el-option label="待删除" value="2" />
            <el-option label="已删除" value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="联系人" prop="contactName">
          <el-input v-model="formData.contactName" placeholder="请输入联系人姓名" />
        </el-form-item>
        <el-form-item label="电话" prop="contactPhone">
          <el-input v-model="formData.contactPhone" placeholder="请输入联系人电话" />
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="formData.address" placeholder="请输入地址" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="formData.description" type="textarea" />
        </el-form-item>
        <el-form-item label="排序" prop="sort">
          <el-input-number v-model="formData.sort" :min="0" :max="999" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog" :loading="btnLoading">确 定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getIedoOrgOrganizationsList,
  createIedoOrgOrganizations,
  updateIedoOrgOrganizations,
  deleteIedoOrgOrganizations,
  findIedoOrgOrganizations
} from '@/plugin/org/api/iedoOrgOrganizations'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { CirclePlus, Edit, Delete } from '@element-plus/icons-vue'
import { filterDict, getDictFunc } from '@/utils/format'
import { useUserStore } from '@/pinia/modules/user'

const treeData = ref([])
const currentOrg = ref(null)
const dialogFormVisible = ref(false)
const type = ref('create')
const formData = ref({
  id: undefined,
  parentId: undefined,
  name: '',
  code: '',
  type: '',
  logo: '',
  status: '1',
  contactName: '',
  contactPhone: '',
  address: '',
  description: '',
  sort: 0,
  createdBy: undefined,
  updatedBy: undefined
})
const orgtypeOptions = ref([])
const RULES = {
  name: [{ required: true, message: '请输入组织名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入组织编码', trigger: 'blur' }],
  type: [{ required: true, message: '请选择组织类型', trigger: 'change' }],
  status: [{ required: true, message: '请选择状态', trigger: 'change' }]
}
const TREE_PROPS = { children: 'children', label: 'name', value: 'id' }
const defaultExpandedKeys = ref([])
const userStore = useUserStore()

const getTableData = async () => {
  const res = await getIedoOrgOrganizationsList({ page: 1, pageSize: 9999 })
  if (res.code === 0) {
    treeData.value = listToTree(res.data.list)
    defaultExpandedKeys.value = treeData.value.map(i => i.id)
  }
}
function listToTree(list, parentId = 0) {
  return list
    .filter(i => i.parentId === parentId)
    .sort((a, b) => (b.sort ?? 0) - (a.sort ?? 0))
    .map(i => ({
      ...i,
      children: listToTree(list, i.id)
    }))
}
const handleNodeClick = (data) => {
  currentOrg.value = data
}
function generateShortCodeGlobal(type, existingCode = null) {
  // 如果是编辑操作且已有 code，则保持原有 code
  if (existingCode) {
    return existingCode;
  }
  
  const prefix = type.charAt(0).toUpperCase();
  // 全部已用code
  const allCodes = [];
  function collectCodes(tree) {
    tree.forEach(item => {
      if (item.type === type && item.code && item.code.startsWith(prefix)) {
        allCodes.push(item.code);
      }
      if (item.children && item.children.length) {
        collectCodes(item.children);
      }
    });
  }
  collectCodes(treeData.value);
  // 提取已用序号
  const usedNums = allCodes
    .map(code => {
      const match = code.match(/^[A-Z](\d{2})$/);
      return match ? parseInt(match[1], 10) : null;
    })
    .filter(num => num !== null);
  let num = 1;
  while (usedNums.includes(num)) num++;
  const numStr = num < 10 ? '0' + num : '' + num;
  return prefix + numStr;
}
function findNodeById(tree, id) {
  for (const node of tree) {
    if (node.id === id) return node
    if (node.children) {
      const found = findNodeById(node.children, id)
      if (found) return found
    }
  }
  return null
}
function generateOrgPath(parentId, code) {
  if (!parentId || parentId === 0) {
    return `/${code}`
  } else {
    const parent = findNodeById(treeData.value, parentId)
    if (!parent) return `/${code}`
    return `${parent.path.replace(/\/$/, '')}/${code}`
  }
}
const initFormData = (parentId = undefined) => {
  return {
    id: undefined,
    parentId: parentId === 0 ? 0 : parentId,
    name: '',
    code: '',
    type: '',
    logo: '',
    status: '1',
    contactName: '',
    contactPhone: '',
    address: '',
    description: '',
    sort: 0,
    createdBy: undefined,
    updatedBy: undefined
  }
}
const addRootOrg = async () => {
  type.value = 'create'
  formData.value = initFormData(0)
  dialogFormVisible.value = true
}
const addChildOrg = async (data) => {
  type.value = 'create'
  formData.value = initFormData(data.id)
  dialogFormVisible.value = true
}
const editOrg = async (data) => {
  const res = await findIedoOrgOrganizations({ id: data.id })
  if (res.code === 0) {
    formData.value = res.data
    type.value = 'update'
    dialogFormVisible.value = true
  }
}
const deleteOrg = (data) => {
  ElMessageBox.confirm('确定要删除该组织吗?', '提示', { type: 'warning' })
    .then(() => deleteIedoOrgOrganizations({ id: data.id }))
    .then(res => { if (res.code === 0) getTableData() })
}
const closeDialog = () => { dialogFormVisible.value = false }
const enterDialog = async () => {
  btnLoading.value = true;
  // 只在创建时生成新的 code
  if (type.value === 'create') {
    formData.value.code = generateShortCodeGlobal(formData.value.type);
  }
  formData.value.path = generateOrgPath(formData.value.parentId, formData.value.code);
  formData.value.level = getLevel(formData.value.parentId);
  if (type.value === 'create') {
    formData.value.createdBy = userStore.userInfo.ID;
    formData.value.updatedBy = userStore.userInfo.ID;
  }
  console.log('最终提交数据:', formData.value); // 调试用
  let res;
  if (type.value === 'create') res = await createIedoOrgOrganizations(formData.value);
  else res = await updateIedoOrgOrganizations(formData.value);
  btnLoading.value = false;
  if (res.code === 0) {
    ElMessage.success('操作成功');
    closeDialog();
    await getTableData();// 刷新树数据
     // 关键修改：重新获取当前组织最新数据
     if (currentOrg.value && currentOrg.value.id) {
      const detailRes = await findIedoOrgOrganizations({ id: currentOrg.value.id });
      if (detailRes.code === 0) {
        currentOrg.value = detailRes.data; // 直接赋值新对象
      }
    }
  }
}
const btnLoading = ref(false)

const getOrgTypeOptions = async () => {
  orgtypeOptions.value = await getDictFunc('orgtype')
}

function getOrgTypeTagType(type) {
  const typeMap = {
    group: 'primary',   // 蓝 集团
    company: 'default',   // 灰蓝 公司
    branch: 'success',   // 绿 分公司
    subsidiary: 'warning',      //  黄 子公司
    department: 'danger',    // 红 部门
    post: 'info'    // 灰 岗位
  }
  return typeMap[type] || 'info'
}

const uploadUrl = import.meta.env.VITE_BASE_API + '/organization/uploadLogo'
function handleLogoSuccess(res) {
  if (res.code === 0) {
    formData.value.logo = res.data
    ElMessage.success('Logo上传成功')
  }
}
function beforeLogoUpload(file) {
  const isImage = file.type.startsWith('image/')
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!')
    return false
  }
  return true
}

function getLevel(parentId) {
  if (!parentId || parentId === 0) return 1;
  const parent = findNodeById(treeData.value, parentId);
  return parent ? (parent.level || 1) + 1 : 1;
}

onMounted(() => {
  getTableData()
  getOrgTypeOptions()
})
</script>

<style scoped>
.org-layout { padding: 20px; background: #f5f7fa; min-height: 100vh; }
.org-main-row { height: 80vh; }
.org-tree-col, .org-detail-col, .org-member-col { height: 100%; }
.org-tree-card, .org-detail-card, .org-member-card { height: 100%; }
.org-tree-header, .org-detail-header, .org-member-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px; }
.custom-tree-node { display: flex; align-items: center; justify-content: space-between; padding-right: 8px; font-size: 14px; }
.node-label { margin-left: 4px; }
/* 操作按钮默认隐藏，悬停节点时显示 */
.custom-tree-node .node-actions {
  opacity: 0;
  transition: opacity 0.2s;
  display: flex;
  gap: 6px; /* 按钮间距加大 */
  margin-left: 8px; /* 与标签拉开距离 */
}
.custom-tree-node:hover .node-actions {
  opacity: 1;
}
.org-detail-card { max-height: 80vh; overflow-y: auto; padding-right: 8px; }
.org-desc :deep(.el-descriptions__label) { white-space: nowrap; min-width: 120px; max-width: 120px; width: 1%; text-align: center; }
</style>
