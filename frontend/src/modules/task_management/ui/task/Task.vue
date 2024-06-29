<script setup lang="ts">
    import type { TaskContract } from '../../domain/contract/taskContract'
    import { useTask } from '../../composable/useTask'
    import StatusIndicator from '@/components/widgets/StatusIndicator.vue'
    import type { Task, TaskStatus } from '../../domain/entity/task'
    import ConfirmPopup from 'primevue/confirmpopup'
    import TaskForm from './Form.vue'
    import Filter from './Filter.vue'
    import { onMounted, ref } from 'vue'
    import { useRouter } from 'vue-router'
    
    const router = useRouter();    
    
    const props = defineProps<{
        TaskRepository : TaskContract,
    }>();

    const taskRecordId = ref(0);
    const formRef = ref();
    const taskStatus = ref<TaskStatus[]>([]);
    
    const {
        loadingRecords,
        tasks,
        filter,
        taskId,
        savingRecord,
        showFilter,
        showDialog,
        totalRecords,
        getAllTasks,
        getTaskStatus
    } = useTask({
        taskRepository: props.TaskRepository
    });

    const save = async () => {
        savingRecord.value = true;
        const result = await formRef.value.save();        
        savingRecord.value = false;        
        if(!result) return;
        showDialog.value = false;
    }

    const openFormDialog = (record: Task | null) => {
        taskRecordId.value = 0;
        if(record) {
            taskRecordId.value = record.id;
            router.push({ params: {id: record.id } });
        }

        showDialog.value = true;
    }

    onMounted(() => {
        const route = router.currentRoute.value;

        if(route.params.id) {
            taskRecordId.value = Number(route.params.id);
            showDialog.value = true;
        }
    });

    const onDialogClose = () => {
        router.push({path: '/tasks'});
    }
    
    getAllTasks();

    getTaskStatus().then((status: TaskStatus[])=>{
        taskStatus.value = status || [];
    });

</script>

<template>
    
    <ContentWrapper @onContentFinish="getAllTasks" :searching="loadingRecords">        

        <template #desktop>
            <div class="card">
        
                <div class="buttons-toolbar">
                    <PButton v-tooltip.bottom="$t('COMMON_BUTTONS.filter')" @click="showFilter = true"  icon="pi pi-search" severity="success" aria-label="Search" />
                    <PButton v-tooltip.bottom="$t('COMMON_BUTTONS.new')" @click="openFormDialog(null)" icon="pi pi-plus" severity="info" aria-label="Task" />
                </div>

                <DataTable :value="tasks" 
                    size="small"
                    scrollable
                    scrollHeight="calc(100vh - 300px)"
                    showGridlines tableStyle="min-width: 50rem" 
                >
                    <template #header>
                        <div class="flex corn">
                            <h1>{{ $t('TASK.task_list')}}</h1>
    
                            <div style="min-width: 40%;">
                                <InputSearch v-model="filter.filters.search" @onChange="getAllTasks" />
                            </div>
                        </div>
                    </template>
        
                    <Column field="title" :header="$t('TASK.title')" style="width: 30%"></Column>
                    <Column field="description" :header="$t('TASK.description')" style="width: 30%"></Column>
                    <Column field="isActive" :header="$t('SECURITY.status')" style="width: 8%; padding-left: 25px;">
                        <template #body="{ data }">
                            <StatusIndicator :style="{ '--bg-color': (data.color) }" :label="$t('TASK.TASK_STATUS.' + (data.status))"></StatusIndicator>
                        </template>
                    </Column>
                    <Column :exportable="false" style="width: 20%">
                        <template #body="{ data }">
                            <div class="grid-actions-container">
                                <!-- <PButton @click="confirmStatusChange($event, data)" class="grid-button-text" icon="pi pi-lock"  v-tooltip.top="'Eliminar'"  text rounded /> -->
                                <PButton @click="openFormDialog(data)" class="grid-button-text" icon="pi pi-pencil" v-tooltip.top="'Editar'" text rounded  />
                            </div>
                        </template>
                    </Column>
                    <template #footer>
                        <PaginationInfo :shown="tasks.length" :total="totalRecords"></PaginationInfo>                        
                    </template>                                   
                </DataTable>     
                <ProgressBar v-if="loadingRecords" mode="indeterminate" style="height: 1px"></ProgressBar>
                <ConfirmPopup></ConfirmPopup>           
            </div>
        </template>

        <template #mobile>
            <div>
                <section class="mobile-input-search">
                    <h2>
                        <i class="pi pi-task" style="font-size: 1.3rem; margin-right: 10px;"></i>
                        {{ $t('TASK.task_list') }}
                    </h2>
                    
                    <div>
                        <InputSearch v-model="filter.filters.search"   @onChange="getAllTasks" />
                    </div>
                </section>

                <section class="mobile-card-wrapper">
                    <MobileGridCard v-for="task in tasks" v-bind:key="task.id">
                        
                        <template #left>
                            <h3>{{ task.title }}</h3>
                            <span>{{ $t('TASK.description')}}: {{ task.description }}</span>
                        </template>
    
                        <template #right>
                            <StatusIndicator :style="{ '--bg-color': (task.color) }" :label="$t('TASK.TASK_STATUS.' + (task.status))"></StatusIndicator>
                        </template>
    
                        <template #actions>
                            <PButton @click="openFormDialog(task)" icon="pi pi-pencil" text />
                            <!-- <PButton @click="confirmStatusChange($event, task)" icon="pi pi-lock" text /> -->
                        </template>
    
                    </MobileGridCard>
                </section>
                <ProgressSpinner v-if="loadingRecords" style="width: 25px; height: 25px; margin: auto" strokeWidth="4" animationDuration="1.5s" />
                <ConfirmPopup></ConfirmPopup>
            </div>
        </template>

        <template #mobile-buttons>            
            <PButton :label="$t('COMMON_BUTTONS.filter')" @click="showFilter = true" text aria-label="Filter">
                <i class="pi pi-filter"></i>
                <span>{{ $t('COMMON_BUTTONS.filter') }}</span>
            </PButton>                               
            <PButton :label="$t('COMMON_BUTTONS.new')" @click="openFormDialog(null)" text aria-label="Nuevo">
                <i class="pi pi-plus"></i>
                <span>{{ $t('COMMON_BUTTONS.new') }}</span>
            </PButton>   
        </template>

        <template #mobile-extra-buttons></template>
    </ContentWrapper>
    
    <PDialog  @hide="onDialogClose" v-model:visible="showDialog" :closeOnEscape="false" modal :header="$t('TASK.' + (taskId === '' ? 'task_registration': 'task_update') )" :style="{ width: '88vw' }">
        <div>
            <TaskForm :taskRepositoy="TaskRepository" :taskId="taskRecordId" ref="formRef" />
        </div>
        <template #footer>            
            <PButton @click="showDialog = false" severity="warning" :label="$t('COMMON_BUTTONS.cancel')" icon="pi pi-times" />
            <PButton @click="save" :label="$t('COMMON_BUTTONS.save')" icon="pi pi-check" autofocus />
        </template>
    </PDialog>

    <Sidebar v-model:visible="showFilter"  position="right">
        <Filter v-model="filter.filters" :taskStatus="taskStatus" @onChage="getAllTasks"></Filter>
    </Sidebar>
</template>

<style lang="scss"></style>