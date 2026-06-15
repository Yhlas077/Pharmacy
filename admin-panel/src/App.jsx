import { Admin, Resource, List, Datagrid, TextField } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';

const dataProvider = simpleRestProvider('http://localhost:8080/api/admin');

const CategoryList = () => (
    <List>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
        </Datagrid>
    </List>
);

export default function App() {
    return (
        <Admin dataProvider={dataProvider}>
            <Resource name="categories" list={CategoryList} />
        </Admin>
    );
}