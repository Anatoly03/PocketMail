<template>
    <div class="view-mails">
        <n-data-table
            remote
            :columns="columns"
            :data="visibleRows"
            :loading="isLoading"
            :pagination="pagination"
            @update:page="openPage"
            @update:filters="updateFilters"
            @update:sorter="updateSorter"
            :render-cell="renderRow"
            :row-key="(row) => row.id"
            :row-props="rowProps"
            flex-height
            :style="{ height: `100%` }"
        >
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
import pb from "../../services/api";
import { h, onMounted, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import type { DataTableColumns, DataTableRowKey, PaginationInfo, PaginationProps } from "naive-ui";
import { NSplit, NDataTable, NIcon, useLoadingBar } from "naive-ui";
import { MailAllRead20Regular } from "@vicons/fluent";
import { FilterState, InternalRowData, TableBaseColumn } from "naive-ui/es/data-table/src/interface";

const loadingBar = useLoadingBar();
const router = useRouter();

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
        width: 200,
        ellipsis: {
            tooltip: true,
        },
    },
    {
        title: "Subject",
        key: "subject",
        sorter: "default",
        ellipsis: {
            tooltip: true,
        },
    },
    {
        title: "Date",
        key: "created",
        width: 90,
        titleAlign: "right",
        align: "right",
        defaultSortOrder: "descend",
        sorter: (row1: RowData, row2: RowData) => row1.sent.getTime() - row2.sent.getTime(),
        render(row: RowData) {
            // if the mail was sent less than few minutes, show "Just now"
            const MINUTE = 60 * 1000; // 1 minute in ms
            if (Date.now() - row.sent.getTime() < 5 * MINUTE) {
                return "Just now";
            }

            // if the mail was sent less than a day ago, show the time only
            const TODAY = 24 * 60 * MINUTE; // 1 day in ms
            if (Date.now() - row.sent.getTime() < TODAY) {
                return row.sent.toLocaleTimeString(undefined, { hour: "2-digit", minute: "2-digit" });
            }

            // if the mail was sent less than 2 days ago, show "Yesterday"
            const YESTERDAY = 2 * TODAY; // 2 days in ms
            if (Date.now() - row.sent.getTime() < YESTERDAY) {
                return "Yesterday";
            }

            // if the mail was sent less than a week ago, show weekday name
            const CUR_WEEK = 7 * TODAY; // 1 week in ms
            if (Date.now() - row.sent.getTime() < CUR_WEEK) {
                return row.sent.toLocaleDateString(undefined, { weekday: "long" });
            }

            // if the mail was sent less than a month ago, show day and month
            const CUR_MONTH = 30 * TODAY; // 1 month in ms
            if (Date.now() - row.sent.getTime() < CUR_MONTH) {
                return row.sent.toLocaleDateString(undefined, { day: "numeric", month: "short" });
            }

            // if the mail was sent less than a year ago, show month and year
            const CUR_YEAR = 365 * TODAY; // 1 year in ms
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
    let new_sort = (sorter.order === "ascend" ? "+" : "-") + sorter.columnKey;
    if (new_sort !== pb_sort.value) {
        pb_sort.value = new_sort;
        query(1);
    }
}

/**
 * Render a cell in the table. This function only overrides `subject` column
 * to make it a link to the mail view page.
 */
const rowProps = (row: RowData) => ({
    style: { cursor: "pointer" },
    onClick: async () => {
        loadingBar.start();
        const res = await router.push(`/mail/${row.id}`);
        if (!res) loadingBar.error();
    },
});

/**
 * On Mount
 */
onMounted(async () => {
    await query(0);
});
</script>

<style lang="scss" scoped>
.view-mails {
    width: 100%;
    height: calc(100% - 16px);
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

<style lang="scss">
.n-data-table .n-data-table__pagination {
    justify-content: center; /* or flex-end for right, flex-start for left */
}
</style>
