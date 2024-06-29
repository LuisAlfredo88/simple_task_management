<script setup lang="ts">
    import type { TaskContract } from '../../domain/contract/taskContract'
    import { useTask } from '../../composable/useTask'

    const props = defineProps<{
		taskRepositoy: TaskContract,
        taskId?: number,
	}>();

    const { 
        taskForm,
		saveTask,
        getTaskById,
        setTaskData,
	} = useTask({
        taskRepository: props.taskRepositoy
    });	

    const loadInitialData = () => {

        if(props?.taskId) {
            getTaskById(props.taskId).then((task) => {
                if(task) {
                    setTaskData(task);
                }
            });
        }
    }
    
    loadInitialData();
    defineExpose({ save: saveTask });
</script>

<template>
    <form ref="formRef">
    <div class="flexbox-grid container">
        <div class="card" style="--min: 50ch">
            <div class="grid">
                <h2>{{  $t('SHARED.basic_information') }}</h2>

				<span class="p-float-label">
					<InputText v-model="taskForm.title" />
                    <label for="title">{{ $t('TASK.title') }}</label>
                </span>

                <span class="p-float-label">
					<InputText v-model="taskForm.description" />
                    <label for="description">{{ $t('TASK.description') }}</label>
                </span>
			</div>
        </div>
    </div>

    </form>
</template>

<style scoped>

</style>