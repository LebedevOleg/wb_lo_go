function GetOrderForm() {
	const [orderId, setOrderId] = React.useState("");
	const [triger, setTriger] = React.useState(false);
	const [orderData, setOrderData] = React.useState([]);
	const press = async () => {
		setOrderData([]);
		await fetch(`http://localhost:3000/orders/${orderId}`)
			.then((res) => res.json())
			.then((result) => {
				setOrderData(Array(result));
				setTriger(true);
			}, []);
	};
	const handlerChangeOrderId = (event) => {
		setOrderId(event.target.value);
	};
	return (
		<div>
			<div>
				<input type="text" onChange={handlerChangeOrderId}></input>
				<button onClick={press}>Получить order по ID</button>
			</div>

			<div>
				Counter: {orderId}
				<br /> Order:{orderData.length} {"  "}
				{(!!orderData.length !== 0 &&
					orderData.map((o) => <OrderItem value={o}></OrderItem>)) ||
					"Данных нет"}
			</div>
		</div>
	);
}

const OrderItem = (props) => {
	return (
		<>
			<table>
				<tr>
					<td>Название поля</td>
					<td>Значение</td>
					<td>Название поля</td>
					<td>Значение</td>
				</tr>
				<tr>
					<td>order_uid</td>
					<td>{props.value.order_uid}</td>
					<td>track_number</td>
					<td>{props.value.track_number}</td>
				</tr>
				<tr>
					<td>customer_id</td>
					<td>{props.value.customer_id}</td>
					<td>delivery_service</td>
					<td>{props.value.delivery_service}</td>
				</tr>
				<tr>
					<td>entry</td>
					<td>{props.value.entry}</td>
					<td>internal_signature</td>
					<td>{props.value.internal_signature}</td>
				</tr>
				<tr>
					<td>date_created</td>
					<td>{props.value.date_created}</td>
					<td>locale</td>
					<td>{props.value.locale}</td>
				</tr>
				<tr>
					<td>oof_shard</td>
					<td>{props.value.oof_shard}</td>
					<td>shardkey</td>
					<td>{props.value.shardkey}</td>
				</tr>
				<tr>
					<td>sm_id</td>
				</tr>
			</table>
		</>
	);
};

ReactDOM.createRoot(document.getElementById("app")).render(<GetOrderForm />);
