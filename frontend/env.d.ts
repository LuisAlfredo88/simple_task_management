/// <reference types="vite/client" />

type GenericKeyValue = {
    id: number | string;
    name: string;
}

type CriteriaFilter = {
	filters: any;
	limit: number;
	skip?: number;
    search?: string;
}

type RequestResponse = {
    status: number;
    data: any;
    headers: any;
}

type GridRecord = {
	records: any[];
	totalRecordsQty: number;
}
