import { useParams } from '@solidjs/router';

const SearchQuery = () => {
  const params = useParams();
  return <div>SearchQuery {JSON.stringify(params.query)}</div>;
};
export default SearchQuery;
