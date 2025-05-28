import { useState, useEffect } from 'react';

import { UserResponse } from '../api/types';
import { useGetUserMutation } from '../api/userApi';

// TODO maybe delete later

const useFetchUsers = (incomingMatches = []) => {
  const [getUser] = useGetUserMutation();
  const [usersData, setUsersData] = useState<UserResponse[]>([]);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const results = await Promise.all(
          incomingMatches.map(async id => {
            try {
              const res = await getUser({}).unwrap();
              return res;
            } catch (err) {
              console.error(`Error fetching user ${id}:`, err);
              return null;
            }
          }),
        );

        setUsersData(results.filter(user => user !== null));
      } catch (err) {
        console.error('Error fetching users:', err);
      }
    };

    if (incomingMatches.length > 0) {
      fetchUsers();
    } else {
      setUsersData([]);
    }
  }, [incomingMatches, getUser]);

  return usersData;
};

export default useFetchUsers;
