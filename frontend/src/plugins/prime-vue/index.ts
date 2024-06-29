import PrimeVue from 'primevue/config'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import ToastService from 'primevue/toastservice'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Row from 'primevue/row'
import Dialog from 'primevue/dialog'
import Toolbar from 'primevue/toolbar'
import Tooltip from 'primevue/tooltip'
import Sidebar from 'primevue/sidebar'
import MobileGridCard from '@/components/widgets/MobileGridCard.vue'
import ContentWrapper from '@/components/wrappers/ContentWrapper.vue'
import PaginationInfo from '@/components/widgets/PaginationInfo.vue'
import InputSearch from '@/components/form/InputSearch.vue'
import Skeleton from 'primevue/skeleton'
import Checkbox from 'primevue/checkbox'
import ConfirmationService from 'primevue/confirmationservice'
import ProgressBar from 'primevue/progressbar'
import ProgressSpinner from 'primevue/progressspinner'
import InputNumber from 'primevue/inputnumber'

const usePrimeVue = (app: any) => {
    app.use(PrimeVue);
    app.use(ToastService);
    app.component('InputText', InputText);
    app.component('PButton', Button);
    app.component('DataTable', DataTable);
    app.component('Column', Column);
    app.component('Row', Row);   
    app.component('Toolbar', Toolbar);
    app.component('PDialog', Dialog);      
    app.component('Sidebar', Sidebar);
    app.component('Skeleton', Skeleton);
    app.component('Checkbox', Checkbox);
    app.component('InputSearch', InputSearch);
    app.component('ContentWrapper', ContentWrapper);          
    app.component('MobileGridCard', MobileGridCard);  
    app.component('ProgressBar', ProgressBar);       
    app.component('ProgressSpinner', ProgressSpinner);  
    app.component('InputNumber', InputNumber);  
    app.component('PaginationInfo', PaginationInfo);           
    
    app.use(ConfirmationService);   
    app.directive('tooltip', Tooltip);
}

export default usePrimeVue;