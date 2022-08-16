import rawLogsOperationsStyles from "@/pages/DataLogs/components/QueryResult/Content/RawLog/RawLogsOperations/index.less";
import { Pagination } from "antd";
import { useModel } from "@@/plugin-model/useModel";
import { useIntl } from "umi";
import { FIRST_PAGE } from "@/config/config";
import { useMemo } from "react";
import { PaneType } from "@/models/datalogs/types";
import HistogramSwitch from "@/pages/DataLogs/components/QueryResult/Content/RawLog/RawLogsOperations/SwitchLeft";

const RawLogsOperations = () => {
  const {
    logCount,
    pageSize,
    currentPage,
    onChangeLogsPage,
    currentLogLibrary,
    doGetLogsAndHighCharts,
    onChangeLogPane,
    logPanesHelper,
    resetLogPaneLogsAndHighCharts,
  } = useModel("dataLogs");
  const { logPanes } = logPanesHelper;
  const i18n = useIntl();

  const oldPane = useMemo(() => {
    if (!currentLogLibrary?.id) return;
    return logPanes[currentLogLibrary?.id.toString()];
  }, [currentLogLibrary?.id, logPanes]);

  const performTime = useMemo(() => {
    return logPanes[currentLogLibrary?.id || 0]?.logs?.cost;
  }, [logPanes]);

  return (
    <div className={rawLogsOperationsStyles.rawLogsOperationsMain}>
      <div className={rawLogsOperationsStyles.operationsBtn}>
        <HistogramSwitch />
      </div>
      {performTime ? (
        <div style={{ flex: 1, textAlign: "right", marginRight: "20px" }}>
          {i18n.formatMessage({ id: "log.perform.time" })}：
          {logPanes[currentLogLibrary?.id || 0]?.logs?.cost}ms
        </div>
      ) : null}
      <div className={rawLogsOperationsStyles.pagination}>
        <Pagination
          size={"small"}
          total={logCount}
          pageSize={pageSize}
          current={currentPage}
          showTotal={(total) => {
            if (!oldPane?.histogramChecked) {
              return false;
            }
            return i18n.formatMessage(
              { id: "log.pagination.total" },
              { total }
            );
          }}
          onChange={(current: number, size: number) => {
            onChangeLogsPage(current, size);
            const params = {
              page: size === pageSize ? current : FIRST_PAGE,
              pageSize: size,
            };
            doGetLogsAndHighCharts(currentLogLibrary?.id as number, {
              isPaging: true,
              reqParams: params,
            })
              .then((res) => {
                if (!res) {
                  resetLogPaneLogsAndHighCharts({
                    ...(oldPane as PaneType),
                    page: size === pageSize ? current : FIRST_PAGE,
                    pageSize: size,
                  });
                } else {
                  const pane: PaneType = {
                    ...(oldPane as PaneType),
                    page: size === pageSize ? current : FIRST_PAGE,
                    pageSize: size,
                    logs: res.logs,
                    highCharts: res.highCharts,
                    logChart: { logs: [] },
                  };
                  onChangeLogPane(pane);
                }
              })
              .catch(() =>
                resetLogPaneLogsAndHighCharts({
                  ...(oldPane as PaneType),
                  page: size === pageSize ? current : FIRST_PAGE,
                  pageSize: size,
                })
              );
          }}
          showSizeChanger
        />
      </div>
    </div>
  );
};
export default RawLogsOperations;
