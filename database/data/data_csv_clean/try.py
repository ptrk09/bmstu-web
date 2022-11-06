# id,name,login,password,role_id


from Cryptodome.Hash import SHA256
import csv


with open('hosts.csv', newline='') as csvr:
    with open('users.csv', 'w', newline='') as csvw:
        reader = csv.reader(csvr, delimiter=',')

        fieldnames = ['id', 'name', 'login', 'password', 'role_id']
        writer = csv.DictWriter(csvw, fieldnames=fieldnames, delimiter=';')    
        
        idx = 0
        for row in reader:
            if idx == 0:
                idx += 1
                continue
            id = row[0]
            name = row[1]
            calculated_host_listings_count = row[2]
            login = name.replace(' ', '_').lower()
            login_bytes = bytes(login, 'utf-8')
            password = SHA256.new(login_bytes).hexdigest()
            role_id = '1'
            
            data = {'id': id, 'name': name, 'login': login, 'password': password, 'role_id': role_id}
            print(str(data))

            writer.writerow(data)
