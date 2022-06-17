import DataAnalysisStyle from "../../index.less";
import { Select, Tooltip } from "antd";
import { useIntl, useModel } from "umi";
import { useEffect, useMemo } from "react";

const ScreeningRow = (props: { style?: any }) => {
  const { style } = props;
  const i18n = useIntl();
  const {
    doGetDatabase,
    instances,
    onChangeCurrentInstances,
    doGetInstance,
    setInstances,
    realTimeTraffic,
    workflow,
  } = useModel("dataAnalysis");
  const { setDatabases, setTables, setNodes, setEdges } = realTimeTraffic;
  const { setIsFold } = workflow;

  useEffect(() => {
    doGetInstance.run().then((res) => setInstances(res?.data ?? []));
  }, []);

  const options = useMemo(() => {
    if (instances.length <= 0) return [];
    return instances.map((item) => ({
      label: (
        <Tooltip
          title={`${item.name}${item.desc && `(${item.desc})`}`}
          placement={"right"}
        >{`${item.name}${item.desc && `(${item.desc})`}`}</Tooltip>
      ),
      value: item.id,
    }));
  }, [instances]);

  return (
    <div className={DataAnalysisStyle.screeningRow} style={style}>
      <Select
        showSearch
        allowClear
        // size="small"
        style={{ width: "278px" }}
        options={options}
        placeholder={i18n.formatMessage({ id: "datasource.draw.selected" })}
        onChange={(iid: number) => {
          setDatabases([]);
          setTables([]);
          setNodes([]);
          setEdges([]);
          setIsFold(false);
          onChangeCurrentInstances(iid);
          if (iid) {
            doGetDatabase
              .run(iid as number)
              .then((res) => setDatabases(res?.data ?? []));
          }
        }}
      />
    </div>
  );
};
export default ScreeningRow;