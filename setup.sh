#!/bin/bash
rm -fr home                                                    # remove all directories
mkdir -p home/john/{work,Downloads,Desktop,Documents,.ssh}     # create john directories
mkdir -p home/jane/{Downloads,Desktop,Documents}               # create jane directories

# reference: https://unix.stackexchange.com/questions/30903/how-to-escape-quotes-in-shell
sudo mkdir home/root                                           # simulate creation of root directory
sudo bash -c "echo $'You can\'t encrypt this' > home/root/root.txt"   # add content in root directory 

# create files in john work directory
cat > home/john/work/Accounts.md << EOF
# Accounts
john:greatercity
admin:fortholder
EOF

cat > home/john/work/Todo.txt << EOF
[ ] Stop using text, switch to markdown
[x] Grow your skillsets
EOF

# reference: https://gist.github.com/monkeywithacupcake/9e0092733302668b2a4adbbfb1d35748
# create an image in john's work folder
create_image () {
  convert -size 200x300 xc:"pink" +repage \
  -size 100x100 -fill white -background None \
  -font NewCenturySchlbk -gravity center caption:"Nice Trick" +repage \
  -gravity Center -composite -strip home/john/work/pix.png
}
create_image

# create ssh key pair for john
ssh-keygen -t rsa -b 2046 -N "" -f home/john/.ssh/id_rsa

# create files in jane work directory
cat > home/jane/Documents/plan.md << EOF
1. Learn python from https://youtu.be/HjGTcasl7jE
2. Create a lab for practice
3. Repeat the process twice in 3days interval
EOF

cat > home/jane/Documents/reward.txt << EOF
- Take yourself out once you finish learning
- Buy yourself a gift on mastery of the content
EOF