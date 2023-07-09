import Checkbox from '@mui/material/Checkbox';
import { Link } from '@mui/material';

function Todo() {
  return (
    <div>
      <div>
        <Checkbox />
        <Link >
          <span>タイトルを書く</span>
        </Link>
      </div>
    </div>
  )
}

export default Todo;
