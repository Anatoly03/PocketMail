<template>
    <div class="view-mails">
        <n-data-table remote :columns="columns" :data="visibleRows" :loading="isLoading" :pagination="pagination" @update:page="openPage" @update:filters="updateFilters" @update:sorter="updateSorter" :row-key="(row) => row.id">
            <template #empty>
                <div class="view-no-mails">
                    <n-icon size="40">
                        <MailAllRead20Regular />
                    </n-icon>
                    No emails.
                </div>
            </template>
        </n-data-table>
    </div>
</template>

<script lang="ts" setup>
import pb from "../services/api";
import { onMounted, reactive, ref } from "vue";
import type { DataTableColumns, DataTableRowKey, PaginationInfo, PaginationProps } from "naive-ui";
import { NSplit, NDataTable, NIcon } from "naive-ui";
import { MailAllRead20Regular } from "@vicons/fluent";
import { FilterState, TableBaseColumn } from "naive-ui/es/data-table/src/interface";

/**
 *
 */
interface RowData {
    id: string;
    sender: string;
    subject: string;
    sent: Date;
}

/**
 * The currently visible rows in the table.
 */
const visibleRows = ref<RowData[]>([]);

/**
 * Is currently fetching new entries?
 */
const isLoading = ref<boolean>(true);

/**
 * Reactive pagination info.
 */
const pagination = reactive<Partial<PaginationProps>>({
    page: 1,
    pageCount: 1,
    pageSize: 12,
    itemCount: 0,
});

const pb_sort = ref<string>("-created");

/**
 * The columns of the table are static.
 */
const columns: DataTableColumns<RowData> = [
    {
        type: "selection",
        disabled(row: RowData) {
            return false;
        },
    },
    {
        title: "Sender",
        key: "sender",
        sorter: "default",
    },
    {
        title: "Subject",
        key: "subject",
        sorter: "default",
    },
    {
        title: "Date",
        key: "created",
        sorter: (row1: RowData, row2: RowData) => row1.sent.getTime() - row2.sent.getTime(),
        render(row: RowData) {
            const MINUTE = 60 * 1000; // 1 minute in ms
            const TODAY = 24 * 60 * MINUTE; // 1 day in ms
            const YESTERDAY = 2 * TODAY; // 2 days in ms
            const CUR_WEEK = 7 * TODAY; // 1 week in ms
            const CUR_MONTH = 30 * TODAY; // 1 month in ms
            const CUR_YEAR = 365 * TODAY; // 1 year in ms

            // if the mail was sent less than few minutes, show "Just now"
            if (Date.now() - row.sent.getTime() < 5 * MINUTE) {
                return "Just now";
            }

            // if the mail was sent less than a day ago, show the time only
            if (Date.now() - row.sent.getTime() < TODAY) {
                return row.sent.toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit" });
            }

            // if the mail was sent less than 2 days ago, show "Yesterday"
            if (Date.now() - row.sent.getTime() < YESTERDAY) {
                return "Yesterday";
            }

            // if the mail was sent less than a week ago, show weekday name
            if (Date.now() - row.sent.getTime() < CUR_WEEK) {
                return row.sent.toLocaleDateString(undefined, { weekday: "long" });
            }

            // if the mail was sent less than a month ago, show day and month
            if (Date.now() - row.sent.getTime() < CUR_MONTH) {
                return row.sent.toLocaleDateString(undefined, { day: "numeric", month: "short" });
            }

            // if the mail was sent less than a year ago, show month and year
            if (Date.now() - row.sent.getTime() < CUR_YEAR) {
                return row.sent.toLocaleDateString(undefined, { month: "short", year: "numeric" });
            }

            return row.sent.toLocaleString();
        },
    },
];

/**
 *
 */
async function query(page: number = pagination.page ?? 1) {
    isLoading.value = true;

    try {
        const res = await pb.collection("mails").getList(page, pagination.pageSize, {
            sort: pb_sort.value,
        });
        console.debug(res);

        visibleRows.value = res.items.map((item: any) => ({
            id: item.id,
            sender: item.sender,
            subject: item.subject,
            sent: new Date(item.created),
        }));

        pagination.page = res.page;
        pagination.pageCount = res.totalPages;
        pagination.itemCount = res.totalItems;
        pagination.pageSize = res.perPage;
    } catch (err) {
        console.error(err);
    }

    isLoading.value = false;
}

/**
 * Open a specific page in the table.
 */
async function openPage(page: number) {
    await query(page);
}

/**
 * Update the filters in the table.
 */
function updateFilters(filterState: FilterState, sourceColumn: TableBaseColumn) {
    console.debug("Filters updated:", filterState, sourceColumn);
}

/**
 * Update the sorter in the table.
 */
function updateSorter(sorter: any) {
    let new_sort = (sorter.order === 'ascend' ? '+' : '-') + sorter.columnKey;
    if (new_sort !== pb_sort.value) {
        pb_sort.value = new_sort;
        query(1);
    }
}

/**
 * On Mount
 */
onMounted(async () => {
    await query(0);
});
</script>

<style lang="scss" scoped>
.view-home {
    width: 100%;
}

.view-no-mails {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 8px;
    color: gray;
}
</style>
