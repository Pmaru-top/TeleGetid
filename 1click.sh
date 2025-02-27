#!/bin/bash

version="1.0.0"
project_name="telegetid"

installation_path="$HOME/.local/bin/$project_name"
version_file="$installation_path/.version"
config_file="$installation_path/config.json"
binary_file="$installation_path/release"
service_file="/etc/systemd/system/$project_name.service"
service_name="$project_name.service"
arch=$(uname -m)

download_latest_release() {
    echo "Downloading file"
    mkdir -p "$installation_path"
    curl https://github.com/pmaru-top/telegetid/releases/download/$version/telegetid_linux_$goarch -L -o $binary_file || exit 1

    # check ELF header to ensure the binary is valid
    if [ "$(head -c 4 $binary_file)" != $'\x7fELF' ]; then
        echo "Failed to download telegetid binary."
        rm -f $binary_file
        exit 1
    fi
    
    echo $version > $version_file
}

install_service() {
    echo "Check the file $binary_file"
    # check binary file
    if [ ! -f "$binary_file" ];then
        echo "Binary file not found!"
        exit 1
    fi

    echo "Setting execute permission for the binary..."
    chmod +rwx $binary_file

    echo "Creating systemd service..."
    sudo bash -c "cat <<EOL > $service_file
[Unit]
Description=$project_name service
After=network.target

[Service]
Environment="PATH=$installation_path"
WorkingDirectory=$installation_path
ExecStart=$binary_file
Restart=always
User=$(whoami)

[Install]
WantedBy=multi-user.target
EOL"

    echo -e "If you want to set it up later\n
    you can press enter ($config_file)"

    read -p "Enter your bot token: " token
    read -p "Enter the proxy address or press enter" proxy

    cat <<EOF > "$config_file"
{
    "token": "$token",
    "proxy": "$proxy"
}
EOF

    isStart=true
    if [ -n "$token" ]; then
        echo "$config_file has been created"
    else
        echo "ok, Please edit the token later in the $config_file and \n start:  sudo systemctl restart $service_name"
        isStart=false
    fi

    echo "Enabling the service..."
    sudo systemctl daemon-reload
    sudo systemctl enable "$service_name"

    if [ "$isStart"=true ]; then
        sudo systemctl restart "$service_name"
    fi

    echo "TeleGetid was successfully installed"
    echo "installation_path: $installation_path"
    echo "service_file_path: $service_file"
}

uninstall_service(){
    sudo systemctl stop $service_name
    sudo systemctl disable $service_name

    rm -rf $installation_path
    sudo rm -f $service_file

    echo "TeleGetid was successfully uninstalled"
}

# check cpu architecture
case "$arch" in
    amd64 | x86_64)
        goarch="amd64"
        ;;
    aarch64 | armv7*)
        goarch="arm"
        ;;
    i386 | i686)
        goarch="386"
        ;;
    *)
        echo "Unsupported architecture: $arch, please install manually."
        exit 1
        ;;
esac

if [ -f "$binary_file" ] || [ -f "$version_file"  ]; then
    echo "Check that $project_name is installed"
    echo -e " [0] Exit \n [1] Reinstall \n [2] Uninstall: "
    read -p "Choose an option: " choose
    case "$choose" in
    0)
        exit 1
        ;;
    1)
        uninstall_service
        download_latest_release
        install_service

        ;;
    2)
        uninstall_service
        ;;

    *)
        exit 1
        ;;
    esac

else
     if [ -f "$version_file" ]; then
        current_version=$(cat "$version_file")
        if [ "$current_version" != "$version" ]; then
            download_latest_release
            install_service
        fi
    else
        download_latest_release
        install_service
    fi
fi