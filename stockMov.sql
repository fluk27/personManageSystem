DECLARE @StartDate DATETIME = null -- '2020-10-01'
        , @EndDate DATETIME = null -- '2020-10-31'


;WITH SI AS (
        SELECT
                S.LotNo
                , S.ItemCode
                , MIN(S.StockId) StockId
                , SUM(IIF(S.ReceiveDate < @StartDate, S.ReceiveQty - ISNULL(S.AdjustQty, 0), 0)) AS PreQty
                , SUM(IIF(S.ReceiveDate >= @StartDate, S.ReceiveQty - ISNULL(S.AdjustQty, 0), 0)) AS ReceiveQty
                , SUM(S.ReceiveQty - ISNULL(S.AdjustQty, 0)) AS BalanceQty
                , SUM(IIF(S.ReceiveDate < @StartDate, S.ReceiveWeight - ISNULL(S.AdjustWeight, 0.0), 0.0)) AS PreWeight
                , SUM(IIF(S.ReceiveDate >= @StartDate, S.ReceiveWeight - ISNULL(S.AdjustWeight, 0.0), 0.0)) AS ReceiveWeight
                , SUM(S.ReceiveWeight - ISNULL(S.AdjustWeight, 0.0)) AS BalanceWeight
        FROM Stock AS S
                WHERE ReceiveDate <= ISNULL(@EndDate, ReceiveDate)
        GROUP BY S.LotNo
                , S.ItemCode
), SO AS (        
        SELECT
                OD.LotNo
                , OD.ItemCode
                , SUM(IIF(D.DispatchDate < @StartDate, OD.ConfirmQty, 0)) AS PreQty
                , SUM(IIF(D.DispatchDate >= @StartDate, OD.ConfirmQty, 0)) AS DispatchQty
                , SUM(OD.ConfirmQty) AS BalanceQty
                , SUM(IIF(D.DispatchDate < @StartDate, OD.ConfirmWeight, 0)) AS PreWeight
                , SUM(IIF(D.DispatchDate >= @StartDate, OD.ConfirmWeight, 0)) AS DispatchWeight
                , SUM(OD.ConfirmWeight) AS BalanceWeight
        FROM dbo.OrderDetail AS OD
        JOIN dbo.[Order] AS O ON O.OrderId = OD.OrderId
        JOIN (
            SELECT OrderId,
                MIN(DispatchDate) AS DispatchDate
            FROM dbo.Dispatch
            WHERE Dispatch.IsActive = 1
            GROUP BY OrderId
        ) AS D ON O.OrderId = D.OrderId
    WHERE O.OrderType=2
                AND DispatchDate <= ISNULL(@EndDate, DispatchDate)
        GROUP BY OD.LotNo
                , OD.ItemCode
), SIO AS (
        SELECT
                S.CustProdCode
                , S.ProductDesc
                , S.ProductLot AS CustomerLotNo
                , SI.ItemCode
                , CS.ServiceName
                , S.ReceiveDate
                , S.LotNo
                , ISNULL(SI.PreQty, 0) - ISNULL(SO.PreQty, 0) AS PreQty
                , ISNULL(SI.PreWeight, 0) - ISNULL(SO.PreWeight, 0) AS PreWeight
                , ISNULL(SI.ReceiveQty, 0) AS ReceiveQty
                , ISNULL(SI.ReceiveWeight, 0.0) AS ReceiveWeight
                , ISNULL(SO.DispatchQty, 0) AS DispatchQty
                , ISNULL(SO.DispatchWeight, 0.0) AS DispatchWeight
                , ISNULL(SI.BalanceQty, 0) - ISNULL(SO.BalanceQty, 0) AS BalanceQty
                , ISNULL(SI.BalanceWeight, 0.0) - ISNULL(SO.BalanceWeight, 0.0) AS BalanceWeight
        FROM SI
                JOIN Stock AS S ON S.StockId = SI.StockId
                JOIN [Order] AS O ON O.OrderId = S.OrderId
                JOIN OrderDetail AS OD ON OD.OrderDetailId = S.OrderDetailId
                LEFT JOIN CustomerService CS ON CS.ServiceId = O.ServiceId
                LEFT JOIN SO ON SI.LotNo = SO.LotNo AND SI.ItemCode = SO.ItemCode
        WHERE ISNULL(SI.PreQty, 0) - ISNULL(SO.PreQty, 0) > 0
                OR ISNULL(SI.ReceiveQty, 0) > 0
                OR ISNULL(SO.DispatchQty, 0) > 0
                OR ISNULL(SI.BalanceQty, 0) - ISNULL(SO.BalanceQty, 0) > 0
)
SELECT *
FROM SIO
-- 818728
-- 465579
-- 139745