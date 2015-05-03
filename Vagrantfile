Vagrant.configure(2) do |config|
  config.vm.box = "wgarcia/centos65-jenkins"

  config.vm.network "forwarded_port", guest: 8080, host: 8888

  config.vm.provider :virtualbox do |box|
    box.name = "centos65-jenkins"
  end
end
